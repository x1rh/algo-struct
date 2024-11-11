package skiplist

import (
    "math/rand"
)

const (
    maxLevel = 32
    pFactor  = 0.25
)

type node struct {
    val  int
    next []*node
}

type Skiplist struct {
    head  *node
    level int
}

func (sl *Skiplist) Search(key int) bool {
    j := sl.head
    for i := sl.level - 1; i >= 0; i-- {
        for j.next[i] != nil && j.next[i].val < key {
            j = j.next[i]
        }
    }
    j = j.next[0]
    return j != nil && j.val == key
}

func (sl *Skiplist) Add(key int) {
    pre := make([]*node, maxLevel)
    for i := range pre {
        pre[i] = sl.head
    }
    for i, j := sl.level-1, sl.head; i >= 0; i-- {
        for j.next[i] != nil && j.next[i].val < key {
            j = j.next[i]
        }
        pre[i] = j
    }

    lv := sl.randomLevel()
    if lv > sl.level {
        sl.level = lv
    }
    newNode := &node{key, make([]*node, lv)} // 注意
    for i, p := range pre[:lv] {
        newNode.next[i] = p.next[i] // new node next = prev node next
        p.next[i] = newNode         // prev node next = new node
    }
}

func (sl *Skiplist) Erase(key int) bool {
    pre := make([]*node, maxLevel)
    j := sl.head
    for i := sl.level - 1; i >= 0; i-- {
        for j.next[i] != nil && j.next[i].val < key {
            j = j.next[i]
        }
        pre[i] = j
    }
    j = j.next[0]
    if j == nil || j.val != key {
        return false
    }
    for i := 0; i < sl.level && pre[i].next[i] == j; i++ {
        pre[i].next[i] = j.next[i]
    }
    for sl.level > 1 && sl.head.next[sl.level-1] == nil {
        sl.level -= 1
    }
    return true
}

func (*Skiplist) randomLevel() (level int) {
    for level = 1; level < maxLevel && rand.Float64() < pFactor; level++ {
    }
    return
}

func NewSkiplist() *Skiplist {
    return &Skiplist{
        head:  &node{-1, make([]*node, maxLevel)},
        level: 0,
    }
}