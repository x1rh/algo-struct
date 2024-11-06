package maths

// Euler 欧拉筛法
// 返回[0, n]范围内的质数数组p
// 返回np数组, np[i] = false 表示i是质数
func Euler(n int) (np []bool, p []int) {
    np = make([]bool, n+1)
    for i:=2; i<=n; i++ {
        if !np[i] {
            p = append(p, i)
        }
        for _, x := range p {
            if x * i > n {
                break 
            }
            np[x*i] = true 
            if i % x == 0 {
                break 
            }
        }
    }
    return 
}
