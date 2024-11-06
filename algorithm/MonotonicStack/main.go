package main

import "fmt"

// 单调增的单调栈，求使得a[l] < a[i] <= a[r]成立的最接近i的l和r
func getlr(a []int) ([]int, []int) {
	n := len(a)
	l := make([]int, n)
	r := make([]int, n)
	s := make([]int, 0, n) // stack
	for i := 0; i < n; i++ {
		l[i] = -1
		r[i] = n
	}

	for i, x := range a {
		for len(s) > 0 && a[s[len(s)-1]] >= x { // 注意等号, 视情况而定
			last := s[len(s)-1]
			s = s[:len(s)-1]
			r[last] = i
		}
		if len(s) > 0 {
			l[i] = s[len(s)-1]
		}
		s = append(s, i)
	}

	return l, r
}

func main() {
	a := []int{1, 2, 3, 4, -1, -2, -3}
	l, r := getlr(a)
	fmt.Println(a)
	fmt.Println(l)
	fmt.Println(r)
}
