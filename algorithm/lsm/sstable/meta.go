package sstable

import (
	"bufio"
	"encoding/binary"
	"errors"
	"os"
)

// .sstable file metadata
type metablock struct {
    dataOffset, dataLen int64 
    indexOffset, indexLen int64
}

// write metadata to the end of file fp 
func (m *metablock) store(fp *os.File) error {
    _, err := fp.Seek(0, 2)    
    if err != nil {
        return err 
    }
    buf := bufio.NewWriter(fp)
    a := []int64{m.dataOffset, m.dataLen, m.indexOffset, m.indexLen}
    for _, x := range a {
        err := binary.Write(buf, binary.BigEndian, x)
        if err != nil {
            return err 
        }
    }
     
    return buf.Flush()
}

// read metadata, read the last 4*8Byte of file fp
func (m *metablock) load(fp *os.File) error { 
    info, err := fp.Stat()
    if err != nil {
        panic(err)
    }
    if info.Size() == 0 {
        return errors.New("invalid empty .sstable, filename=" + info.Name())
    }

    _, err = fp.Seek(-32, 2)
    if err != nil { 
        return err 
    } 

    buf := bufio.NewReader(fp) 
    a := make([]int64, 4)
    for i := range a {        
        err = binary.Read(buf, binary.BigEndian, &a[i]) 
        if err != nil {
            return err 
        } 
    }
    m.dataOffset = a[0]
    m.dataLen = a[1]
    m.indexOffset = a[2]
    m.indexLen = a[3]
    
    return nil 
}
