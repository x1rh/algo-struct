package main

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

type SlowDataStruct []int

func (s SlowDataStruct) Add(l, r, v int) {
	for i := l - 1; i < r; i++ {
		s[i] += v
	}
}

func (s SlowDataStruct) Set(l, r, v int) {
	for i := l - 1; i < r; i++ {
		s[i] = v
	}
}

func (s SlowDataStruct) Query(l, r int) int {
	var res int
	for i := l - 1; i < r; i++ {
		res += s[i]
	}
	return res
}

func NewSlowDataStruct(a []int) SlowDataStruct {
	sds := make(SlowDataStruct, len(a))
	copy(sds, a)
	return sds
}

func TestSegmentTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	testCnt := 150
	maxv := 10000
	maxn := 100000
	maxop := 100000

	t.Logf("test case total=%d", testCnt)
	for kase := 1; kase <= testCnt; kase++ {
		n := rand.Intn(maxn) + 1
		opCnt := rand.Intn(maxop) + 1

		a := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = rand.Intn(maxv)
		}

		t.Logf("case=%d, n=%d, opCnt=%d", kase, n, opCnt)

		seg := InitSegmentTree(a)
		sds := NewSlowDataStruct(a)

		for i := 0; i < opCnt; i++ {
			op := rand.Intn(3)
			ql, qr := 1, n
			l := rand.Intn(n)
			r := rand.Intn(n)
			v := rand.Intn(maxv)
			f := rand.Intn(2)
			if f == 1 {
				v = -v
			}
			if r > l {
				l, r = r, l
			}

			l++
			r++

			if l < 1 || r > n {
				t.Fatal("invalid l, r")
			}

			if op == 0 {
				seg.Add(1, 1, n, l, r, v)
				sds.Add(l, r, v)
			} else if op == 1 {
				seg.Set(1, 1, n, l, r, v)
				sds.Set(l, r, v)
			} else {
				ql = l
				qr = r
			}
			if sds.Query(ql, qr) != seg.Query(1, 1, n, ql, qr) {
				t.Fatal("fail!!!")
			}
		}

		a = nil
		seg = nil
		sds = nil
		runtime.GC()
	}
}
