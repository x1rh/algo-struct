package maths

// Eratosthenes 埃式筛
// 返回长度为n+1的布尔数组np[], 覆盖区间[0, n]
// np[i] = false 表示数i是质数 
func Eratosthenes(n int) []bool {
	np := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if !np[i] {
			for j := i * i; j <= n; j += i { // j=i*i是正确的, i, 2i, 3i, ... (i-1)*i 都已经被筛过了
				np[j] = true
			}
		}
	}
	return np
}
