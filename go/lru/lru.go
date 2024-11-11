package lru  

import (
	"container/list"
)

// entry 同时记录key 和 val，在通过链表元素删除哈希表元素时，需要用到key 
type entry struct {
    key, val int     
}

type LRUCache struct {
    size int 
    capacity int 
    cache map[int]*list.Element
    l *list.List          
}

func (lru *LRUCache) Get(key int) int {
    v := lru.cache[key]
    if v == nil {
        return -1 
    } else {
        lru.l.MoveToFront(v)
        return v.Value.(entry).val
    } 
}

func (lru *LRUCache) Put(key, val int) {
    v := lru.cache[key]
    e := entry{key, val}
    if v != nil {
        v.Value = e   // key 对应的 val可能发生改变 
        lru.l.MoveToFront(v)
    } else {
        lru.cache[key] = lru.l.PushFront(e)
        if lru.l.Len() > lru.capacity {
            delete(lru.cache, lru.l.Remove(lru.l.Back()).(entry).key)
        }
    }
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        size: 0,
        capacity: capacity,
        cache: make(map[int]*list.Element, capacity),
        l: list.New(), 
    }
}

