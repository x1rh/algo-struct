package mathx

import (
	"math/rand"
	"testing"
)

func TestEratosthenes(t *testing.T) {
	maxn := 1000000
	n := rand.Intn(maxn) + 1
	t.Logf("TestEratosthenes in [2, %d]", n)
	a := Eratosthenes(n)
	b := BruteForce(n)
	for i := 2; i <= n; i++ {
		if a[i] != b[i] {
			t.Fatal("a[i] != b[i]")
		}
	}
}
