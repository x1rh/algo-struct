package lockfreequeue

import (
	"errors"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestListQueue1(t *testing.T) {
	rand.Seed(time.Now().Unix())
	maxn := 1000
	a := make([]int, maxn)
	q := NewListQueue()
	for i := 0; i < maxn; i++ {
		x := rand.Intn(maxn)
		a[i] = x
		q.EnQueue(x)
	}

	for i := 0; i < maxn; i++ {
		x, _ := q.DeQueue()
		if a[i] != x {
			t.Fatal("not equal")
		}
	}
}

func TestListQueue2(t *testing.T) {
    q := NewListQueue()
    var wg sync.WaitGroup
    n := 3 
    m := 10000000
    var r, w int64 
    for i:=0; i<n; i++ {
        wg.Add(1)
        go func() {
            for i:=0; i<m; i++ {
                op := rand.Intn(2)
                if op == 0 {
                    q.EnQueue(rand.Int())
                    atomic.AddInt64(&w, 1)
                } else {
                    if _, err := q.DeQueue(); err == nil {
                        atomic.AddInt64(&r, 1)
                    } else if !errors.Is(err, ErrEmpty) {
                        panic(err)
                    }
                }
            } 
            wg.Done()
        }() 
    }
    wg.Wait()
    for {
        if _, err := q.DeQueue(); err != nil {
            if errors.Is(err, ErrEmpty) {
                break 
            }
            panic(err)
        } else {
            atomic.AddInt64(&r, 1)            
        }
    }
    if r != w {
        t.Fatalf("not equal, %d != %d\n", w, r)
    } 
} 

func TestEmpty(t *testing.T) {
    defer func() {
        if err := recover(); err != nil {
            if !errors.Is(err.(error), ErrEmpty) {
                panic(err)
            }
        }
    }() 
    
    q := NewListQueue()
    q.EnQueue(1)
    x, _ := q.DeQueue()
    if x != 1 {
        t.Fatal("not equal")
    }
    
    if _, err := q.DeQueue(); !errors.Is(err, ErrEmpty) {
        panic(err)
    }
}
