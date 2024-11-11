package closeOpenInterval

func LowerBound(a []int, x int) int {
	l, r := 0, len(a)
	for l < r {
		m := l + (r-l)/2
		if a[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}
	// 答案一定是l，因为它永远指向正确的位置(存在或不存在)
	// 结束时 l>=r, 且只能是l == r
	return l
}

func UpperBound(a []int, x int) int {
	l, r := 0, len(a)
	for l < r {
		m := l + (r-l)/2
		if a[m] <= x {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
