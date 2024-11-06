package sstable

import (
    "testing"
)

func TestIndex(t *testing.T) {
}

func TestCal(t *testing.T) {
    a := []struct{a, b int}{
        {0, 0}, 
        {1, 1},
        {2, 2},
        {3, 2},
        {4, 2},
        {5, 3},
    }
    idx := newIndexBlock()
    for _, x := range a {
        if idx.cal(x.a) != x.b {
            t.Log(idx.cal(x.a), x.b)
            t.Fatal("not equal")
        }
    }
}
