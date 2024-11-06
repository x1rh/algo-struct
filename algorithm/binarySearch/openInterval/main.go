package openInterval

func LowerBound(a []int, x int) int {
	l, r := -1, len(a)
	for l+1 < r {
		m := l + (r-l)/2
		if a[m] < x {
			l = m
		} else {
			r = m
		}
	}
	// 退出时, l+1 == r
	// 所以返回值为 l + 1 或者 r
	return l + 1
}

func UpperBound(a []int, x int) int {
	l, r := -1, len(a)
	for l+1 < r {
		m := l + (r-l)/2
		if a[m] <= x {
			l = m
		} else {
			r = m
		}
	}
	return r
}
