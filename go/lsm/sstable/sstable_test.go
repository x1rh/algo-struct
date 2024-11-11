package sstable

import (
	"log"
	"os"
	"testing"

	"lsm/command"
	"lsm/memtable"
)


var sstable_test_kv map[string]string = map[string]string{
    "1": "1", 
    "2": "2",
    "222": "222",
    "221": "221",
    "3": "3",
}

var sstable_test_file = "./test.sst"

var sstable_test_not_exist = []string{"11", "22", "2222", "2212", "33", "0", "hello", "test"}

func TestSSTable(t *testing.T) {
    path := sstable_test_file
    
    m1 := memtable.New()
    for k, v := range sstable_test_kv {
        m1.Put(k, command.Command{Op: command.OpPut, Key: k, Val: v})
    }
    
    sst1, err := New(path)
    if err != nil {
        panic(err)
    }
    if err := sst1.Store(m1); err != nil {
        panic(err)
    }
    sst1.Close()
        
    m2, err := getMemtableFromFile(path)
    if err != nil {
        panic(err)
    }

    if !compareTwoMemtable(m1, m2) {
        t.Fatal("not equal")
    }
}


func TestSSTableGet(t *testing.T) { 
    path := sstable_test_file 
    sst, err := New(path, LoadSSTable())
    if err != nil {
        panic(err)
    }

    for k, v1 := range sstable_test_kv {
        v2, ok := sst.Get(k)
        if !ok { 
            t.Fatalf("key = %s not found\n", k)
        }
        if v1 != v2 {
            t.Fatalf("key = %s, v1=%s != v2=%s", k, v1, v2)
        }
    }

    for _, k := range sstable_test_not_exist {
        _, ok := sst.Get(k)
        if ok {
            t.Fatalf("found key = %s\n", k)
        }
    }

}

func compareTwoMemtable(m1, m2 *memtable.MemTable) bool {
    if m1.Size() != m2.Size() {
        log.Printf("different size: %d != %d\n", m1.Size(), m2.Size())
        return false 
    }
    
    n := m1.Size()
    keys1 := m1.Keys()
    keys2 := m2.Keys() 

    for i:=0; i<n; i++ {
        if keys1[i] != keys2[i] {
            log.Printf("different key: %s != %s\n", keys1[i], keys2[i])
            return false 
        }
    }

    vals1 := m1.Values()
    vals2 := m2.Values() 
    for i:=0; i<n; i++ {
        lv := vals1[i].(command.Command)
        rv := vals2[i].(command.Command) 
        if lv != rv {   // notice ==  
            log.Printf("different val: %+v != %+v\n", vals1[i], vals2[i])
            return false 
        }
    }
    
    return true 
}


// LoadMemtableFromFile load a memtable from file, just for testing 
func getMemtableFromFile(path string) (*memtable.MemTable, error) {
    sst, err := New(path, LoadSSTable())
    if err != nil {
        return nil, err 
    } 

    m := memtable.New()
    for _, v := range sst.index.Values() {
        seginfo := v.(indexElem) 
        data, err := sst.loadDataSegment(seginfo.Offset, seginfo.Len)
        if err != nil {
            return nil, err 
        }
        for i := range data {
            m.Put(data[i].Key, data[i])
        } 
    }

    sst.Close()
    
    return m, nil  
}

// clear test file 
func TestClearAll(t *testing.T) {
    if err := os.Remove(sstable_test_file); err != nil {
        panic(err)
    }
}
