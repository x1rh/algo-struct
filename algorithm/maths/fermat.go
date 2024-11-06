package maths

import "math/rand"

const (
    testCnt = 50
)

// fermat 费马素性测试(随机型算法), 判断n是不是素数
// a<n, n是素数(或者a和n互质), 则 a^(n-1) ≡ 1 (mod n) 
// 时间复杂度：O(k log n), k = testCnt 
func Fermat(n int) bool {
    for i:=0; i<testCnt; i++ {
        a := rand.Intn(n-1) + 1
        if ModPow(a, n-1, n) != 1 {
            return false 
        }
    }
    return true 
}
