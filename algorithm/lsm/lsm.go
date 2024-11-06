package lsm

import (
	"container/list"
	"errors"
	"fmt"
	"lsm/command"
	"lsm/memtable"
	"lsm/sstable"
	"lsm/wal"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
    DEFAULT_DATA_DIR = "./datadir"
    SSTABLE_EXT = ".sst"
    WAL_EXT = ".wal"
    // 32 MB
    DEFAULT_THRESHOLD = 32 * 1024 * 1024             
)

// log structured merge tree
type LSMTree struct {
    locker *sync.Mutex            // guard for creating new memtable 
    mutable *memtable.MemTable    // RedBlackTreeMap[string]Command
    immutable *memtable.MemTable  // 存储还未持久化的memtable 
    sstables *list.List             // 存储已经持久化了的sstable
    
    datadir string                // 数据文件夹    
    
    threshold int 
    size int                      // size += len(key) + len(val)
    
    filename string            //  {$filename}.sst, {$filename}.wal   
    wal *wal.Log               // write-ahead log 
}

type LSMTreeOption func(*LSMTree) 

func DataDir(datadir string) LSMTreeOption {
    return func(lsm *LSMTree) {
        lsm.datadir = datadir  
    } 
}

func Threshold(threshold int) LSMTreeOption {
    return func(lsm *LSMTree) {
        lsm.threshold = threshold
    }
}


func NewLSMTree(opts ...LSMTreeOption) (*LSMTree, error) {
    lsm := &LSMTree{
        locker: &sync.Mutex{},
        mutable: memtable.New(),
        sstables: list.New(),
        datadir: DEFAULT_DATA_DIR,
        filename: nextFilename(), 
        threshold: DEFAULT_THRESHOLD,
    }

    for _, opt := range opts {
        opt(lsm)
    }
    
    if info, err := os.Stat(lsm.datadir); errors.Is(err, os.ErrNotExist) {
        if err := os.MkdirAll(lsm.datadir, 0777); err != nil {
            return nil, err 
        }
    } else if !info.IsDir() {
        return nil, errors.New("datadir is a file, not dir")
    } 
     
    lsm.load()
   
    if lsm.wal == nil {
        wal, err := wal.New(filepath.Join(lsm.datadir, lsm.filename + WAL_EXT))
        if err == nil {
            lsm.wal = wal 
        } else {
            return nil, err 
        }
    }
     
    return lsm, nil 
}

// LSM 在工作前 必须把之前全部的wal文件处理
// loadSSTable load each .sstable file from datadir 
// 区分两种恢复情形
// 1. lsm中还未达到阈值就发生异常的异常
//      - 此时将wal文件恢复到lsm的memtable，删除wal文件
//      - 生成新的wal文件
// 2. lsm中还未进行持久化或者持久化进行到一半的sstable，
//      - 此时将wal文件恢复到sstable文件中，删除wal文件。
// bug: 触发重写wal时，过程中如果宕机了，那么可能出现问题，需要一套机制来辨别上次宕机时还未持久化的wal文件和重写该wal文件产生的新wal文件
func (lsm *LSMTree) load() error {
    var f []os.FileInfo 
    if err := filepath.Walk(lsm.datadir, func(path string, info os.FileInfo, err error) error {
        f = append(f, info)
        return nil 
    }); err != nil {
        return err 
    }

    sort.Slice(f, func(i, j int) bool {
        return f[i].Name() > f[j].Name()
    })

    // 需要保证文件遍历顺序从晚到早
    for _, info := range f {
        if info.IsDir() {
            continue 
        }
        path := filepath.Join(lsm.datadir, info.Name())
        pre := strings.Split(filepath.Base(info.Name()), ".")[0]
        ext := filepath.Ext(info.Name()) 
        var sst *sstable.SSTable
        switch ext { 
        case SSTABLE_EXT: 
            walfile := filepath.Join(lsm.datadir, pre + WAL_EXT)
            if _, err := os.Stat(walfile); errors.Is(err, os.ErrNotExist) {
                fmt.Println("case 1")
                sst, err = sstable.New(path, sstable.LoadSSTable())           // 持久化成功的SSTable (只存在.sst文件)
                if err != nil {  
                    return err 
                }
                lsm.sstables.PushBack(sst)
            } else { 
                fmt.Println("case 2")
                w, err := wal.New(path)
                if err != nil {
                    return err 
                }
                m := memtable.New()                                           // 持久化中途宕机的SSTable (同时存在相同名字的.wal文件和.sst文件)
                if _, err := w.Restore(m); err != nil {
                    return err 
                }
                sst, err = sstable.New(path)
                if err != nil {
                    return err 
                }
                if err := sst.Store(m); err != nil {
                    return err 
                }
                if err := w.Delete(); err != nil {
                    return err 
                }
                lsm.sstables.PushBack(sst) 
            }
            
        case WAL_EXT:
            w, err := wal.New(path)
            if err != nil {
                return err 
            }
            tblfile := filepath.Join(lsm.datadir, pre + SSTABLE_EXT)
            if _, err := os.Stat(tblfile); errors.Is(err, os.ErrNotExist) {    // 宕机后恢复MemTable  (只存在.wal文件)
                sz1, err := w.Restore(lsm.mutable)
                if err != nil {
                    return err 
                }
                sz2 := getMemtableSize(lsm.mutable)
                if sz2 >= lsm.threshold {
                    fmt.Println("case 3-1")
                    if err := lsm.store(); err != nil {
                        return err 
                    }
                } else if sz1 >= lsm.threshold {
                    fmt.Println("case 3-2")
                    if err := lsm.rewriteWal(); err != nil {
                        return err 
                    }
                    if err := w.Delete(); err != nil {
                        return err 
                    }
                } else { 
                    fmt.Println("case 3-3")
                    lsm.filename = pre
                    fmt.Println("prefix: ", pre)
                    lsm.wal = w  
                }
                fmt.Println("case 3")
            } else {
                fmt.Println("case 4")
                m := memtable.New()                                            // 持久化中途宕机的SSTable (同时存在相同名字的.wal文件和.sst文件) 
                if _, err := w.Restore(m); err != nil {
                    return err 
                }
                sst, err = sstable.New(path)
                if err != nil {
                    return err 
                }
                if err := sst.Store(m); err != nil {
                    return err 
                }
                if err := w.Delete(); err != nil {
                    return err 
                }
                lsm.sstables.PushBack(sst)
            }
        } 
    }
    return nil 
}

// store current memtable to disk 
// make current mutable memtable to immutable memtable, and create a new mutable memtable  
// 没有考虑并发（虽然我想）
func (lsm *LSMTree) store() error {
    lsm.locker.Lock() 
    defer lsm.locker.Unlock() 

    if lsm.immutable != nil {
        return errors.New("immutable not nil")
    }

    if lsm.size < lsm.threshold {      // 重新判断
        return nil 
    }
    
    lsm.wal.Close()
    lsm.wal = nil 

    // 保持和wal文件名相同
    oldname := lsm.filename
    lsm.filename = nextFilename() 
    oldwalfile := filepath.Join(lsm.datadir, oldname + WAL_EXT)
    oldtblfile := filepath.Join(lsm.datadir, oldname + SSTABLE_EXT) 
    newwalfile := filepath.Join(lsm.datadir, lsm.filename + WAL_EXT)

    // 通过先创建sstable文件，用来区分两种异常恢复
    if _, err := os.Create(oldtblfile); err != nil {   
        return err
    }

    // 创建新的wal文件
    if err := lsm.nextwal(newwalfile); err != nil {
        return err 
    }
 
    lsm.immutable = lsm.mutable
    lsm.mutable = memtable.New() 
    lsm.size = 0 
       
    sst, err := sstable.New(oldtblfile)
    if err != nil {
        return err 
    }
    if err := sst.Store(lsm.immutable); err != nil {           
        return err  
    }   
    if err := os.Remove(oldwalfile); err != nil {
        return err 
    }
    lsm.sstables.PushBack(sst)
    lsm.immutable = nil 
    return nil 
}


func (lsm *LSMTree) Put(key, val string) {
    c := command.Command{Op: command.OpPut, Key: key, Val: val}
    if err := lsm.wal.Write(c); err != nil {
        panic(err)
    }
    lsm.mutable.Put(key, c)
    lsm.size += c.Size() 
    if lsm.size >= lsm.threshold {
        if err := lsm.store(); err != nil {
            panic(err)
        }
    }
}

func (lsm *LSMTree) Get(key string) (string, bool){
    if v, ok := lsm.mutable.Get(key); ok {
        c := v.(command.Command)
        if c.Op == command.OpDel {
            return "", false 
        }
        return c.Val, true 
    } 

    if lsm.immutable != nil {
        if v, ok := lsm.immutable.Get(key); ok {
            c := v.(command.Command)    
            if c.Op == command.OpDel {
                return "", false 
            }
            return c.Val, true 
        }
    }
    if lsm.sstables != nil && lsm.sstables.Len() > 0 {
        for it:=lsm.sstables.Front(); it != nil ; it=it.Next() {
            sst := it.Value.(*sstable.SSTable)
            if val, ok := sst.Get(key); ok {
                return val, ok 
            }
        }        
    }
    return "", false 
}

func (*LSMTree) IntervalQuery(k1, k2 string) []command.Command {
    return nil 
}

func (lsm *LSMTree) Del(key string) {
    c := command.Command{Op: command.OpPut, Key: key}
    if err := lsm.wal.Write(c); err != nil {
        panic(err)
    }
    if v, ok := lsm.mutable.Get(key); ok {
        c := v.(command.Command)
        if c.Op != command.OpDel {
            lsm.mutable.Remove(key)
        }
    } else {
        lsm.mutable.Put(key, c)
        lsm.size += c.Size() 
        if lsm.size >= lsm.threshold {
            if err := lsm.store(); err != nil {
                panic(err)
            }
        }
    }
}


func (*LSMTree) Compact() {
}

func (lsm *LSMTree) Close() {
    lsm.mutable = nil 
    lsm.immutable = nil 
    lsm.wal.Close()
    for i:=lsm.sstables.Front(); i != nil; i=i.Next() {
        i.Value.(*sstable.SSTable).Close()
    }
}

// close and delete prev wal and create a new one 
func (lsm *LSMTree) nextwal(path string) error {
    // lsm.wal.Close()
    if nw, err := wal.New(path); err != nil {
        return err 
    } else {
        lsm.wal = nw 
    }
    return nil 
}

// 重复对相同key操作，导致wal文件膨胀，但memtable大小未超过threshold
// 此时选择重写wal文件到另一个wal文件, 重写完成后再替换
func (lsm *LSMTree) rewriteWal() error {
    filename := nextFilename()
    newwalfile := filepath.Join(lsm.datadir, filename + WAL_EXT)
    w, err := wal.New(newwalfile)
    if err != nil {
        return err 
    }
    if err := w.WriteTable(lsm.mutable); err != nil {
        return err 
    }
    lsm.filename = filename
    lsm.wal = w
    return nil 
}

func nextFilename() string {
    return strconv.Itoa(int(time.Now().UnixNano()))
}


// 重写.wal文件，仅发生在LSMTree.load()时
// 具体情形为：重复对一个key进行操作，wal文件不断膨胀，但


func getMemtableSize(m *memtable.MemTable) int {
    size := 0 
    for _, v := range m.Values() {
        c := v.(command.Command)
        size += c.Size()
    }
    return size 
}

