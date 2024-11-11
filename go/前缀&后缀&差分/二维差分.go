package main

import "fmt"

type Diff struct {
	d [][]int
	m, n int
}

func NewDiff(m, n int) *Diff {
	if m <=0 || n <=0 {
		return nil
	}

	d := make([][]int, m+1)
	for i:=0; i<=m; i++ {
		d[i] = make([]int, n+1)
	}

	return &Diff {d: d, m: m, n:n}
}

func (p *Diff) Add(r1, c1, r2, c2, v int) {
	r2 += 1
	c2 += 1
	p.d[r1][c1] += v
	p.d[r1][c2] -= v
	p.d[r2][c1] -= v
	p.d[r2][c2] += v
}

func (p *Diff) GetMatrix() [][]int {
	r := make([][]int, p.m)
	for i:=0; i<p.m; i++ {
		r[i] = make([]int, p.n)
		for j:=0; j<p.n; j++ {
			if i==0 && j==0 {
				r[i][j] = 0
			} else if i == 0 {
				r[i][j] = r[i][j-1]
			} else if j == 0 {
				r[i][j] = r[i-1][j]
			} else {
				r[i][j] = r[i-1][j] + r[i][j-1] - r[i-1][j-1]
			}
			r[i][j] += p.d[i][j]
		}
	}
	return r
}

func main() {
	d := NewDiff(3, 3)
	d.Add(1, 1, 2, 2, 1)
	d.Add(0, 0, 1, 1, 1)

	m := d.GetMatrix()
	for _, r := range m {
		fmt.Println(r)
	}
}