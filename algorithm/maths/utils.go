package maths

import (
    "math/big"
)

func abs(x int) int {
    if x < 0 {
        return -x  
    }
    return x 
}

func max(x, y int) int {
    if x > y {
        return x 
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x 
    }
    return y
}

// ModMul  return x * y % p 
// 要求返回值在64位内
func BigIntModMul(x, y, p int) int {
    bx := big.NewInt(int64(x))
    by := big.NewInt(int64(y))
    bp := big.NewInt(int64(p))
    res := big.NewInt(0).Mod(big.NewInt(0).Mul(bx, by), bp)
    return int(res.Int64())
}

