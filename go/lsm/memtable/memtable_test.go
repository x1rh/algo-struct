package memtable

import (
    "testing"
    "lsm/command"
)

var memtable_test_kv map[string]string = map[string]string{
    "1": "1", 
    "2": "2",
    "222": "222",
    "221": "221",
    "3": "3",
}

func TestMemTable(t *testing.T) {
    m := New()
    for k, v := range memtable_test_kv {
        m.Put(k, command.Command{Op: command.OpPut, Key: k, Val: v})
    }

    for k, v := range memtable_test_kv {
        x, ok := m.Get(k)
        if !ok {
            t.Fatal("k not found")
        }
        c, ok := x.(command.Command)
        if !ok {
            t.Fatal("x is not a Command")
        }
        if c.Key != k || c.Val != v {
            t.Fatal("not equal")
        }
    } 
    m.Clear()
}

