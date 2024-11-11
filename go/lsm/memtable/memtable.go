package memtable

import (
    "github.com/emirpasic/gods/trees/redblacktree"
)

// memtable 代表当前内存中的键值对表，通常要求O(log N)的操作效率
// 这里选用底层为红黑树的TreeMap
// trick, 为了方便替换底层数据结构, 这里使用继承。可以替换其他数据结构，例如 SkipList
// 也可以额外写个interface，whatever
type MemTable struct {
    redblacktree.Tree 
}

func New() *MemTable {
    m := new(MemTable)
    m.Tree = *redblacktree.NewWithStringComparator()
    return m 
}
