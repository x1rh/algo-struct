package closeInterval

func LowerBound(a []int, x int) int {
	l, r := 0, len(a)-1 // 注意
	for l <= r {        // 保证[l, r]不为空
		m := l + (r-l)/2
		if a[m] < x {
			// 更新l=m+1后意味着l-1指向的是小于key的区域; 而l永远指在了正确的位置
			l = m + 1
		} else {
			r = m - 1
		}
	}
	// l或者r+1
	return l
}

// 如果key大于所有的数组元素，那么l=len(a)
// 如果key小于所有的数组元素，那么r=-1, l=0
// 所以处理返回结果时，先判断l是否数组下标越界，<len(a)，然后判断a[l]是否等于key

func UpperBound(a []int, x int) int {
	l, r := 0, len(a)-1
	for l <= r {
		m := l + (r-l)/2
		if a[m] <= x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	// l 或 r+1
	return l
}
