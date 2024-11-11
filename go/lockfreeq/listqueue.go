package lockfreequeue

import (
	"errors"
	// "fmt"
	"sync/atomic"
	"unsafe"
)

var ErrEmpty = errors.New("empty queue")

type listPointer struct {
    cnt uint                // cnt表示ptr指向的listNode的引用计数
    ptr *listNode 
}

type listNode struct {
	val  int
	next *listPointer
}

type ListQueue struct {
	head, tail *listPointer
}

func NewListQueue() *ListQueue {
	q := &ListQueue{}
    n := &listNode{next: new(listPointer)}
    p := &listPointer{ptr: n} 
	q.head = p
	q.tail = p
	return q
}

// EnQueue 两个关键操作，1）插入新值 2）更新tail指针
func (q *ListQueue) EnQueue(val int) {
	node := &listNode{
		val: val,
        next: new(listPointer),
	}

	var tail, p *listPointer 

    for {
        tail := q.tail  
        next := tail.ptr.next   

        if tail == q.tail {
            if next.ptr == nil {
                p = &listPointer{ptr: node, cnt: next.cnt + 1}
                if atomic.CompareAndSwapPointer(
                    (*unsafe.Pointer)(unsafe.Pointer(&q.tail.ptr.next)),
                    unsafe.Pointer(next),
                    unsafe.Pointer(p),
                ) {
                    break 
                }
            } else {
                p = &listPointer{ptr: next.ptr, cnt: tail.cnt + 1}       // tail.cnt + 1 ? 
                atomic.CompareAndSwapPointer(
                    (*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
                    unsafe.Pointer(tail),
                    unsafe.Pointer(next),
                )
            }
        }
    }
    atomic.CompareAndSwapPointer(
        (*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
        unsafe.Pointer(tail),
        unsafe.Pointer(p), 
    ) 
}

// DeQueue
// 将head指向真正队头节点前一个节点，
func (q *ListQueue) DeQueue() (res int, err error) {
    var head *listPointer 
    for {
        head = q.head 
        tail := q.tail  
        next := head.ptr.next
        if head == q.head {
            if head.ptr == tail.ptr {
                if next.ptr == nil {
                    return 0, ErrEmpty
                } else {            
                    // tail has changed, fetch it 
                    p := &listPointer{ptr: next.ptr, cnt: tail.cnt + 1} 
                    atomic.CompareAndSwapPointer(
                        (*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
                        unsafe.Pointer(tail),
                        unsafe.Pointer(p))
                }
            } else {  
                p := &listPointer{ptr: next.ptr, cnt: head.cnt + 1}
                if atomic.CompareAndSwapPointer(
                    (*unsafe.Pointer)(unsafe.Pointer(&q.head)),
                    unsafe.Pointer(head),
                    unsafe.Pointer(p)) {
                        res = next.ptr.val
                        break 
                }
            }
        }
    }
    return   
}
