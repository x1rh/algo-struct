package main

import "fmt"

type Diff struct {
	d []int
}

func NewDiff(v []int) *Diff {
	if len(v) == 0 {
		return nil
	}

	r := &Diff {
		d: make([]int, len(v)),
	}

	r.d[0] = v[0]
	for i:=1; i<len(v); i++ {
		r.d[i] = v[i] - v[i-1]
	}

	return r
}

// GetArray 获取数组
func (p *Diff) GetArray() []int {
	n := len(p.d)
	res := make([]int, n)
	res[0] = p.d[0] 
	for i:=1; i<n; i++ {
		res[i] = res[i-1] + res[i]
	}

	return res
}

func (p *Diff) Add(l, r, v int) {
	p.d[l] += v
	if r+1 < len(p.d) {
		p.d[r+1] -= v
	}
}


func main() {
	a := []int{1, 2, 3, 4, 5}

	d := NewDiff(a)
	d.Add(0, 4, 1)

	fmt.Println(d.GetArray())
}