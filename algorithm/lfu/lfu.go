package lfu

import (
	treeset "github.com/emirpasic/gods/sets/treeset"
)

type entry struct {
	cnt, time, key, val int
}

type LFUCache struct {
	capacity, time int
	cache map[int]*entry
	set *treeset.Set
}

func cmp(a, b interface{}) int {
	l := a.(entry)
	r := b.(entry)
	if l.cnt < r.cnt {
		return -1
	} else if l.cnt == r.cnt {
		if l.time < r.time {
			return -1
		} else if l.time == r.time {
			return 0
		} else {
			return 1
		}
	} else {
		return 1
	}
}

func (lfu *LFUCache) Get(key int) int {
	v, ok := lfu.cache[key]
	if !ok {
		return -1
	} else {
		lfu.time += 1
		e := entry{v.cnt+1, lfu.time, v.key, v.val}
		lfu.set.Remove(*v)
		lfu.set.Add(e)
		lfu.cache[key] = &e
		return e.val
	}
}

func (lfu *LFUCache) Put(key, val int) {
	var e entry
	lfu.time += 1
	v, ok := lfu.cache[key]
	if !ok {
		if lfu.capacity == len(lfu.cache) && lfu.capacity != 0 {
			it := lfu.set.Iterator()
			it.Next()
			o := it.Value().(entry)
			lfu.set.Remove(o)
			delete(lfu.cache, o.key)
		}
		e = entry{1, lfu.time, key, val}
	} else {
		lfu.set.Remove(*v)
		e = entry{v.cnt+1, lfu.time, key, val}
	}
	lfu.cache[key] = &e
	lfu.set.Add(e)
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		time: 0,
		cache: make(map[int]*entry),
		set: treeset.NewWith(cmp),
	}
}