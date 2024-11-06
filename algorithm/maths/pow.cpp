//go:build ignore
#include <iostream>

// 快速幂-朴素版
int fastpow(int x, int n) {
    int res = 1;
    while(n > 0) {
        if(n & 1) res *= x;
        x *= x;
        n >>= 1;
    }
    return res;
}

//快速幂-取模
int fast_mod_pow(int x, int n, int mod){
	int res = 1;
	while(n > 0) {
		if (n & 1) res = res * x % mod;
		x = x * x % mod;
		n >>= 1;
	}
	return res;
}
