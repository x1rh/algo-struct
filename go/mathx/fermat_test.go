package mathx

import (
	"math/rand"
	"testing"
	"time"
)

func TestFermat(t *testing.T) {
	maxn := 100000000
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(maxn) + 1
	t.Logf("TestFermat int [2, %d]", n)
	for i := 2; i <= n; i++ {
		if IsPrime(n) != Fermat(n) {
			t.Fatalf("isPrime(%d) != fermat(%d)", n, n)
		}
	}
	t.Logf("TestFermat(%d) pass", n)
}
