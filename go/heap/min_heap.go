package heap

import (
	"fmt"
)

func build(a []int) {
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		down(a, i)
	}
}

func down(a []int, i int) bool {
	p := i // parent
	s := 0 // son
	n := len(a)
	if n == 0 {
		return false
	}
	for {
		l := 2*p + 1 // left  son
		r := 2*p + 2 // right son
		s = l
		if l >= n || l < 0 { // l < 0 after int overflow
			break
		}
		if r < n && a[r] < a[l] {
			s = r
		}

		if a[s] >= a[p] {
			break
		}

		a[p], a[s] = a[s], a[p]
		p = s
	}

	return p > i
}

func up(a []int, j int) {
	for {
		i := (j - 1) / 2
		if i == j || a[i] <= a[j] { // NOTE:
			break
		}
		a[i], a[j] = a[j], a[i]
		j = i
	}
}

func push(a []int, idx, val int) {
	a[idx] = val
	up(a, len(a)-1) // NOTE:
}

func pop(a *[]int) int {
	r := (*a)[0]
	(*a)[0] = (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1] // NOTE:
	down(*a, 0)           // NOTE:

	return r
}

func main() {
	a := []int{1, 2, 3, 8, 7, 6, 5, 4, 9, 10, 12, 11}

	build(a)
	for len(a) > 0 {
		x := pop(&a)
		fmt.Println(x)
	}

}
