package maths

// BruteForce 暴力求np数组
// np[i] = false 表示i是质数
// 主要用来对拍测试
func BruteForce(n int) []bool {
	np := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				np[i] = true
				break
			}
		}
	}
	return np
}
