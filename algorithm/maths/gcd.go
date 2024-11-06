package maths

// gcd 返回a和b的最大公因数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b 
}
