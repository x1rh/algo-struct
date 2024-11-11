package mathx

// lcm 返回a和b的最小公倍数
func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}
