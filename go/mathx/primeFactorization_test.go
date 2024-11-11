package mathx

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// pfsc -> PrimeFactorizationSpecialCases
var pfsc []int = []int{
	12,
	1 * 2 * 3 * 5 * 7 * 11 * 13,
	1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13,
	1836514918,
	7757701789,
	8715857146,
}

func TestPrimeFactorizationBF(t *testing.T) {
	for _, x := range pfsc {
		t.Log(PrimeFactorizationBF(x))
	}
}

func TestPrimeFactorization(t *testing.T) {
	for _, x := range pfsc {
		t.Log(PrimeFactorization(x))
	}
}

func TestSpecialCaseEqual(t *testing.T) {
	for _, x := range pfsc {
		map1 := PrimeFactorizationBF(x)
		map2 := PrimeFactorization(x)
		t.Logf("test case x=%d", x)
		s1 := fmt.Sprintf("%v", map1)
		s2 := fmt.Sprintf("%v", map2)
		if s1 != s2 {
			t.Fatalf("x=%d\n%s != %s", x, s1, s2)
		}
	}
}

func TestEqual(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n := 10000
	maxn := 10000000000
	for i := 0; i < n; i++ {
		x := rand.Intn(maxn) + 1
		map1 := PrimeFactorizationBF(x)
		map2 := PrimeFactorization(x)
		t.Logf("test case %d/%d, x=%d", i+1, n, x)
		s1 := fmt.Sprintf("%v", map1)
		s2 := fmt.Sprintf("%v", map2)
		if s1 != s2 {
			t.Fatalf("x=%d\n%s != %s", x, s1, s2)
		}
	}
}

// func BenchmarkPrimeFactorization(b *testing.B) {
//     for i:=1; i<=b.N; i++ {
//         m := PrimeFactorization(i)
//         b.Log(m)
//     }
// }
