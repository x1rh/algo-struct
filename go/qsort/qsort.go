package main

import "fmt"

func QuickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	l, r, k := lo, hi, a[lo]

	for l < r {
		for l < r && a[r] >= k {
			r -= 1
		}
		a[l] = a[r]

		for l < r && a[l] <= k {
			l += 1
		}
		a[r] = a[l]
	}

	a[l] = k
	QuickSort(a, lo, l-1)
	QuickSort(a, l+1, hi)
}


func main() {
	a := []int{5, 4, 3, 2, 1, 6, 7, 8, 12, 11, 10}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}