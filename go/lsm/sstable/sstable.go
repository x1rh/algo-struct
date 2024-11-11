package sstable

import (
	"bufio"
	"encoding/json"
	"errors"
	"lsm/command"
	"lsm/memtable"
	"math"
	"os"
	"sort"
)

// sort string table - SSTable
// 给SSTable建立一个简单的索引，每segsize个键分为一段，用一个TreeMap记录每一段的在文件中的偏移和长度
// 显然最佳分块个数是sqrt(n), n是键的个数
type SSTable struct {
    fp *os.File    
    path string               // full path 
    meta *metablock
    index *indexblock         //  
    segsize int 
}

func New(path string, opts ...sstableOption) (*SSTable, error){
    fp, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
    if err != nil {
        return nil, err 
    }    
    sst := &SSTable{
        path: path,
        fp: fp, 
        index: newIndexBlock(),
        meta: &metablock{},
    }
    for _, opt := range opts {
        if err := opt(sst); err != nil {
            return nil, err
        }
    }  
    
    return sst, nil 
}

type sstableOption func(*SSTable) error 


// load get sstable index from .sstable file
// load index and metadata, do not load all data  
func LoadSSTable() sstableOption {
    return func(sst *SSTable) error {
        if err := sst.meta.load(sst.fp); err != nil {
            return err 
        }
        if err := sst.index.load(sst.fp, sst.meta.indexOffset, sst.meta.indexLen); err != nil {
            return err 
        }
        return nil
    }
}

// Get get k v from current sstable 
func (sst *SSTable) Get(key string) (string, bool) {
    node, ok := sst.index.floor(key)
    if ok {
        for node != nil {
            el := node.Value.(indexElem)
            data, err := sst.loadDataSegment(el.Offset, el.Len)
            if err != nil {
                panic(err)
            }
            if data[len(data)-1].Key < key {
                node = node.Right
                continue 
            }
            j := sort.Search(len(data), func(i int) bool {
                return data[i].Key >= key  
            })
            if j != len(data) && data[j].Key == key {
                if data[j].Op == command.OpDel {
                    return "", false 
                }
                return data[j].Val, true 
            } 
            return "", false     
        } 
    }
    return "", false 
}

// store() write sstable to file 
// content of .sstable file:ffff
// | 0              ~                len |
// | data block | index block | metadata |
// split data block to several segment, index block tell each segment's offset and len  
func (sst *SSTable) Store(table *memtable.MemTable) error { 
    n := table.Size()

    if n == 0 {
        sst.meta = &metablock{}
        return sst.meta.store(sst.fp)  
    } 
    
    sz := int(math.Ceil(math.Sqrt(float64(n))))  // calculate segment size 
    seg := make([]command.Command, sz)           // store each segment of data block  
    var idx []indexElem                          // store all index elems in order 
    offset := 0 
    buf := bufio.NewWriter(sst.fp)

    // write data block to buf   
    for i, v := range table.Values() {           
        seg[i%sz] = v.(command.Command)
        if (i+1) % sz == 0 || i==n-1 {
            if i == n - 1 {
                seg = seg[:i%sz+1]
            }
            dump, err := json.Marshal(seg)
            if err != nil {
                return err 
            }
            sst.index.Put(seg[0].Key, indexElem{Key: seg[0].Key, Offset: offset, Len: len(dump)})
            idx = append(idx, indexElem{seg[0].Key, offset, len(dump)})
            buf.Write(dump)
            offset += len(dump)
        }
    } 

    // write index block to buf  
    dump, err := json.Marshal(idx)
    if err != nil {
        return err 
    }
    buf.Write(dump)
    

    // flush data block and index block to file  
    sst.meta.dataOffset = 0
    sst.meta.dataLen = int64(offset)
    sst.meta.indexOffset = int64(offset)
    sst.meta.indexLen = int64(len(dump))
    if err := buf.Flush(); err != nil {
        return err
    }

    // write and flush metadata  
    return sst.meta.store(sst.fp) 
}

// read data segment from sstable data block 
// data block contains several data segment
func (sst *SSTable) loadDataSegment(Offset, Len int) ([]command.Command, error) {
    if offset, err := sst.fp.Seek(int64(Offset), 0); err != nil {
        return nil, err 
    } else if offset != int64(Offset) {
        return nil, errors.New("offset != Offset")
    } 

    dump := make([]byte, Len)
    if n, err := sst.fp.Read(dump); err != nil {
        return nil, err 
    } else if n != Len {
        return nil, errors.New("n != Len")
    }

    var data []command.Command 
    if err := json.Unmarshal(dump, &data); err != nil {
        return nil, err 
    }
    return data, nil 
}


func (sst *SSTable) Close() {
    sst.fp.Close()
}

