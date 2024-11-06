package main

import "fmt"

// [l, r] = [1, n]

type SegmentTree []struct {
	sumv, addv, setv int
	addf, setf       bool // flag
}

func (t SegmentTree) pushUp(p int) {
	t[p].sumv = t[p<<1].sumv + t[p<<1|1].sumv
}

func (t SegmentTree) pushDown(p, Len int) {
	o := &t[p]
	ll := Len - Len/2 // 左区间长度
	rl := Len / 2     // 右区间长度
	if o.setf {
		t.doSet(p<<1, ll, o.setv)
		t.doSet(p<<1|1, rl, o.setv)
		o.setf = false
	}
	if o.addf {
		t.doAdd(p<<1, ll, o.addv)
		t.doAdd(p<<1|1, rl, o.addv)
		o.addv = 0 
		o.addf = false
	}
}

func (t SegmentTree) doSet(p, Len, setv int) {
	t[p].setv = setv
	t[p].setf = true
	t[p].addf = false
	t[p].addv = 0
	t[p].sumv = Len * setv
}

func (t SegmentTree) doAdd(p, Len, addv int) {
	t[p].addv += addv
	t[p].addf = true
	t[p].sumv += Len * addv
}

func (t SegmentTree) Add(p, l, r, ql, qr, v int) {
	if ql <= l && r <= qr {
		t.doAdd(p, r-l+1, v)
	} else {
		m := (l + r) / 2
		t.pushDown(p, r-l+1)
		if ql <= m {
			t.Add(p<<1, l, m, ql, qr, v)
		}
		if qr > m {
			t.Add(p<<1|1, m+1, r, ql, qr, v)
		}
		t.pushUp(p)
	}
}

func (t SegmentTree) Set(p, l, r, ql, qr, v int) {
	if ql <= l && r <= qr {
		t.doSet(p, r-l+1, v)
	} else {
		m := (l + r) / 2
		t.pushDown(p, r-l+1)
		if ql <= m {
			t.Set(p<<1, l, m, ql, qr, v)
		}
		if qr > m {
			t.Set(p<<1|1, m+1, r, ql, qr, v)
		}
		t.pushUp(p)
	}
}

func (t SegmentTree) Query(p, l, r, ql, qr int) int {
	if ql <= l && r <= qr {
		return t[p].sumv
	} else {
		m, res := (l+r)/2, 0
		t.pushDown(p, r-l+1)
		if ql <= m {
			res += t.Query(p<<1, l, m, ql, qr)
		}
		if qr > m {
			res += t.Query(p<<1|1, m+1, r, ql, qr)
		}
		t.pushUp(p)
		return res
	}
}

func (t SegmentTree) Build(a []int, p, l, r int) {
	if l == r {
		t[p].sumv = a[l-1]
	} else {
		m := (l + r) / 2
		t.Build(a, p<<1, l, m)
		t.Build(a, p<<1|1, m+1, r)
		t.pushUp(p)
	}
}

func InitSegmentTree(a []int) SegmentTree {
	tree := make(SegmentTree, len(a)*4)
	tree.Build(a, 1, 1, len(a))
	return tree
}

func NewSegmentTree(n int) SegmentTree {
	return make(SegmentTree, n*4)
}

func main() {
	a := []int{3, 2, 1, 4, 5, 6}
	tree := InitSegmentTree(a)
	fmt.Println(tree.Query(1, 1, len(a), 1, len(a)))
	tree.Add(1, 1, len(a), 1, len(a), 1)
	fmt.Println(tree.Query(1, 1, len(a), 1, len(a)))
	tree.Add(1, 1, len(a), 1, 1, 1)
	fmt.Println(tree.Query(1, 1, len(a), 1, len(a)))
	tree.Set(1, 1, len(a), 1, 3, 1)
	fmt.Println(tree.Query(1, 1, len(a), 1, len(a)))
	tree.Set(1, 1, len(a), 1, len(a), 0)
	fmt.Println(tree.Query(1, 1, len(a), 1, len(a)))

}
