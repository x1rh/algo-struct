package wal

import (
	"bufio"
	"encoding/json"
	"io"
	"lsm/command"
	"lsm/memtable"
	"os"
)

// write-ahead log
type Log struct {
    fp *os.File
    path string 
}

func New(walfilepath string) (*Log, error) {
    fp, err := os.OpenFile(walfilepath, os.O_RDWR|os.O_CREATE, 0777)
    if err != nil {
        return nil, err  
    }
    return &Log{fp: fp, path: walfilepath}, nil 
}

// 数据格式:  opt key val\n 
func (wal *Log) Write(c command.Command) error {
    if _, err := wal.fp.Seek(0, 2); err != nil {
        return err 
    }
    dump, err := json.Marshal(c)
    if err != nil {
        return err 
    }
    if _, err = wal.fp.Write(dump); err != nil { 
        return err 
    }
    if _, err = wal.fp.Write([]byte{'\n'}); err != nil {
        return err 
    }
    return wal.fp.Sync() 
}


// restore MemTable from .wal file
func (wal *Log) Restore(m *memtable.MemTable) (int, error) {
    wal.fp.Seek(0, 0)
    buf := bufio.NewReader(wal.fp)
    size := 0 
    for {
        line, err := buf.ReadBytes('\n')
        if err == io.EOF {
            break 
        }
        if err != nil {
            return 0, err 
        } 
        c := new(command.Command)
        if err := json.Unmarshal(line, c); err != nil {
            return 0, err 
        }
        size += c.Size() 
        switch c.Op {
        case command.OpPut:
            m.Put(c.Key, *c) 
        case command.OpDel:
            if _, ok := m.Get(c.Key); ok {
                m.Remove(c.Key)
            } else {
                m.Put(c.Key, *c)
            }
        }
    } 
    return size, nil 
}

// clear the content of the write-ahead log 
func (wal *Log) clear() error {
    if err := wal.fp.Truncate(0); err != nil {
        return err 
    }
    if err := wal.fp.Sync(); err != nil {
        return err 
    }
    return nil 
}

func (wal *Log) Delete() error { 
    if err := wal.fp.Close(); err != nil {
        return err 
    }
    wal.fp = nil 
    if err := os.Remove(wal.path); err != nil {
        return err
    } 
    return nil 
}

func (wal *Log) Close() {
    if wal == nil {
        println("fuck you ")
    }
    if wal.fp != nil {
        _ = wal.fp.Close()
    }
}


func (wal *Log) Name() string {
    return wal.path
}

// 将整个MemTable写到文件，从文件开头开始 
func (wal *Log) WriteTable(m *memtable.MemTable) error {
    if _, err := wal.fp.Seek(0, 0); err != nil {
        return err 
    }
    w := bufio.NewWriter(wal.fp)
    for _, v := range m.Values() {
        c := v.(command.Command)
        dump, err := json.Marshal(c)
        if _, err = w.Write(dump); err != nil {
            return err 
        }
        if _, err = w.Write([]byte{'\n'}); err != nil {
            return err 
        }
    } 
    return w.Flush()
}
