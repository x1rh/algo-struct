package mathx

// PrimeFactorizationBF 对n进行质因数分解 (暴力版)
// 例如 n = 1*2*3*4*5*6*7*8*9*10*11*13
func PrimeFactorizationBF(n int) map[int]int {
	f := make(map[int]int)
	for i := 2; n != 1; i++ {
		if n%i == 0 {
			n /= i
			f[i] += 1
			i--
		}
	}
	return f
}

// PrimeFactorization 利用PollardRho 算法进行质因数分解
// todo: 记忆化搜索
func PrimeFactorization(n int) map[int]int {
	f := make(map[int]int) // 记录质因数及其个数
	var r func(int)
	r = func(x int) {
		if MillerRabin(x) { // 必须保证素性判断100%正确
			f[x] += 1
		} else {
			d := PollardRho(x)
			if d > x {
				panic(d)
			}
			r(d)
			r(x / d)
		}
	}
	r(n)
	return f
}
