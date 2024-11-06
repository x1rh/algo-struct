package maths

import (
	"math/big"
	"math/rand"
	"time"
)

// PollardRho 是一个能快速找到大整数n的一个非1、非自身的因子的算法
// 时间复杂度: O(n^(1/4) log n)
func PollardRho(n int) int {
    if n == 4 {
        return 2 
    }
    if MillerRabin(n) {
        return n 
    }

    bn := big.NewInt(int64(n))
    rand.Seed(time.Now().UnixNano())
    f := func(x, c int) int {    // 注意处理溢出 
        bx := big.NewInt(int64(x))
        bc := big.NewInt(int64(c))
        r  := big.NewInt(0)
        r = r.Mul(bx, bx).Add(r, bc).Mod(r, bn)
        return int(r.Int64()) 
    } 
    for {
        c := rand.Intn(n-1) + 1 
        t, r, p, q := 0, 0, 1, 0 
        for {
            for i:=0; i<128; i++ {  // 令固定距离C=128
                t = f(t, c)
                r = f(f(r,c),c)
                if t == r {
                    break 
                }        
                q = BigIntModMul(p, abs(t-r), n)
                if q == 0 {
                    break 
                }
                p = q 
            }
            d := gcd(p, n)
            if d > 1 {
                return d 
            }
            if t == r {
                break 
            }
        }
    }
}


// MaxPrimeFactor 利用PollardRho()求数n的最大质因数
func MaxPrimeFactor(n int) int {
    g := make(map[int]int)      // 记忆化搜索
    max := func(x, y int) int {
        if x > y {
            return x 
        }
        return y 
    }
    get := func(x int) int {
        if v, ok := g[x]; ok {
            return v 
        }
        factor, res := PollardRho(n), 0
        if factor == 1 {
            res = x 
        } else {
            res = max(MaxPrimeFactor(factor), MaxPrimeFactor(x / factor))
        }
        g[x] = res 
        return res 
    }
    return get(n)
}

