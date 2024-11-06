package main

import "fmt"

type SumRange struct {
	prefix []int
	suffix []int
}

func NewSumRange(v []int) *SumRange {
	ret := &SumRange {
		prefix: make([]int, len(v) + 1),
		suffix: make([]int, len(v) + 1),
	}

	for i, x := range v {
		ret.prefix[i+1] = ret.prefix[i] + x
	}

	for i:=len(v)-1; i>=0; i-- {
		ret.suffix[i] = ret.suffix[i+1] + v[i]
	}

	return ret
}

func (s *SumRange) GetPrefix(l, r int) int {
	return s.prefix[r+1] - s.prefix[l]
}

func (s *SumRange) GetSuffix(l, r int) int {
	return s.suffix[l] - s.suffix[r+1]
}


func main() {
	a := []int{1, 2, 3, 4, 5}

	sr := NewSumRange(a)

	fmt.Println(sr.GetPrefix(1, 2))
	fmt.Println(sr.GetPrefix(0, 4))
	fmt.Println(sr.GetSuffix(3, 4))
}