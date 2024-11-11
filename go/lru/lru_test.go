package lru 

import (
    "testing"
)

func TestLRUCache(t *testing.T) {
    lru := NewLRUCache(10) 
    lru.Put(1, 2)
    if lru.Get(1) != 2 {
        t.Fatal("Get(1) != 2")
    }
    lru.Put(1, 3)
    if lru.Get(1) != 3 {
        t.Fatal("Get(1) != 3") 
    }
    
}
