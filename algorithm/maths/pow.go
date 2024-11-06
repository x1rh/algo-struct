package maths

import "math/big"

// ModPow return a^n % p
func ModPow(a, n, p int) int {
    x := 1 
    for n > 0 {
        if n % 2 == 1 {
            x = x * a % p 
        }
        a = a * a % p 
        n /= 2 
    }
    return x 
}

// BigIntModPow 
// 要求返回值在64位内, 但是中途计算值超过64位
func BigIntModPow(a, n, p int) int {
    x := big.NewInt(1)
    ba := big.NewInt(int64(a))
    bp := big.NewInt(int64(p))
    for n > 0 {
        if n % 2 == 1 {
            x = x.Mod(x.Mul(x, ba), bp) 
        }
        ba = ba.Mul(ba, ba).Mod(ba, bp)
        n /= 2 
    }
    return int(x.Int64())
}

// Pow return a^n 
func Pow(a, n int) int {
    x := 1 
    for n > 0 {
        if n % 2 == 1 {
            x = x * a 
        }
        a = a * a 
        n /=  2 
    }
    return x 
}
