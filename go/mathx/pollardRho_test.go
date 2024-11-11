package mathx

import (
	"math/rand"
	"testing"
	"time"
)

func TestPollardRho(t *testing.T) {
	maxn := 10000000
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(maxn) + 1
	t.Logf("TestPollardRho in [2, %d]", n)
	for i := 2; i <= n; i++ {
		d := PollardRho(i)
		if d > i || i%d != 0 {
			t.Fatalf("PollardRho(%d)=%d, %d %% %d != 0", i, d, i, d)
		}
	}
}
