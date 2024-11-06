package main

import "fmt"

// 哈哈，像归并排序，又像快排序，但是太好记了，看一眼就记住了。来自ChatGPT
func qsort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	var l, r []int
	pivot := a[0]
	for _, x := range a[1:] {
		if x < pivot {
			l = append(l, x)
		} else {
			r = append(r, x)
		}
	}
	l = qsort(l)
	r = qsort(r)

	l = append(append(l, pivot), r...)
	return l
}

func main() {
	a := []int{5, 4, 3, 2, 1, 6, 7, 8, 12, 11, 10}

	b := qsort(a)

	for _, x := range b {
		fmt.Printf("%v ", x)
	}
	fmt.Println()
}