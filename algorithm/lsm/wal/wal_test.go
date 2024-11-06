package wal

import (
	"lsm/command"
	"lsm/memtable"
	"os"
	"testing"
)

var test_commands = []command.Command{
    {Op:command.OpDel, Key: "111", Val: "111"},
    {Op: command.OpPut, Key: "222", Val: "222"},
    {Op: command.OpPut, Key: "1", Val: "1"}, 
    {Op: command.OpDel, Key: "2", Val: "2"},
    {Op: command.OpPut, Key: "2", Val: "2"},
    {Op: command.OpDel, Key: "2", Val: "2"},
    {Op: command.OpPut, Key: "2", Val: "2"},
    {Op: command.OpPut, Key: "3", Val: "3"},
    {Op: command.OpPut, Key: "221", Val: "221"},
    {Op: command.OpDel, Key: "221", Val: "221"},
}

var wal_test_file = "./test.wal"

func TestWal(t *testing.T) {
    path := wal_test_file
    w, err := New(path)
    if err != nil {
        panic(err)
    }
    for _, c := range test_commands {
        w.Write(c)
    }

    m := memtable.New()
    _, err = w.Restore(m)    
    if err != nil {
        panic(err)
    }
}



func TestClearAll(t *testing.T) {
    if err := os.Remove(wal_test_file); err != nil {
        panic(err)
    }
}

