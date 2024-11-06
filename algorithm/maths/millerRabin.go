package maths

// MillerRabin 素性判定，判定n是不是素数
// 时间复杂度: O(k log n) ?, 其中k = 7 
func MillerRabin(n int) bool {
    if n < 3 {
        return n == 2 // 特判1, 2 
    }   
    if n % 2 == 0 {   // 特判偶数
        return false 
    }
    magic := []int{2, 325, 9375, 28178, 450775, 9780504, 1795265022}  // 2^64 以内不会出错 
    d := n - 1 
    r := 0 
    for d % 2 == 0 {
        d /= 2 
        r += 1 
    }

    for _, a := range magic {
        v := BigIntModPow(a, d, n)   // a^d % n 
        if v <= 1 || v == n - 1 {
            continue 
        }
        for i:=0; i < r; i++ {
            v = BigIntModMul(v, v, n)
            if v == n - 1 && i != r - 1 {
                v = 1 
                break 
            }
            if v == 1 {
                return false 
            }
        }
        if v != 1 {
            return false 
        }     
    }
    return true 
}
