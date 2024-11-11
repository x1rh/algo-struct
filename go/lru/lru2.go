package lru

import (
	"fmt"
)

// 题目： https://www.nowcoder.com/practice/5dfded165916435d9defb053c63f1e84

type Node struct {
	entry      entry
	prev, next *Node
}

type List struct {
	head, tail *Node
	size       int
}

func (l *List) Front() *Node {
	return l.head
}

func (l *List) PushFront(e entry) {
	node := &Node{entry: e, next: l.head}

	if l.head != nil {
		l.head.prev = node
	}

	l.head = node

	if l.tail == nil {
		l.tail = l.head
	}
}

func (l *List) MoveToFront(node *Node) {
	if node == l.head {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node == l.tail {
		l.tail = node.prev
	}
	l.head.prev = node
	node.next = l.head
	l.head = node
}

func (l *List) PopBack() *Node {
	tail := l.tail
	l.tail = l.tail.prev
	tail.prev = nil
	return tail
}

func (l *List) print() {
	for t := l.head; t != nil; t = t.next {
		fmt.Print(t.entry, " ")
	}
	fmt.Println()
}

type entry struct {
	key, val int
}

type Solution struct {
	size, capacity int
	cache          map[int]*Node
	l              *List
}

func Constructor(capacity int) Solution {
	return Solution{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*Node),
		l:        &List{},
	}
}

func (this *Solution) get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.l.MoveToFront(v)
		return v.entry.val
	} else {
		return -1
	}
}

func (this *Solution) set(key int, value int) {
	if v, ok := this.cache[key]; ok {
		if v.entry.val != value {
			v.entry.val = value
		}
		this.l.MoveToFront(v)
	} else {
		this.l.PushFront(entry{key, value})
		this.cache[key] = this.l.Front()
		if len(this.cache) > this.capacity {
			delete(this.cache, this.l.PopBack().entry.key)
		}
	}
}
