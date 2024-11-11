package mathx

import (
	"math/rand"
	"testing"
	"time"
)

var millerRabinCases []int = []int{
	9433310687,
	4319560489,
	4580128549,
}

// todo: 更完整的测试 && 利用多核
func TestMillerRabin(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	maxn := 10000000
	n := rand.Intn(maxn) + 1
	t.Logf("TestMillerRabin in [2, %d]", n)
	for i := 2; i <= n; i++ {
		if IsPrime(i) != MillerRabin(i) {
			t.Fatalf("IsPrime(%d) != MillerRabin(%d)", i, i)
		}
	}
}

func TestMillerRabinSpecialCases(t *testing.T) {
	for _, x := range millerRabinCases {
		if IsPrime(x) != MillerRabin(x) {
			t.Fatalf("IsPrime(%d) != MillerRabin(%d)", x, x)
		}
	}
}
