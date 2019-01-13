package subarray

import (
	"flag"
	"math/rand"
	"os"
	"testing"
)

var a []int
var n uint

func TestMain(m *testing.M) {
	flag.UintVar(&n, "n", 1000, "Number of inputs")
	flag.Parse()
	a = []int{}
	for i := uint(0); i < n; i++ {
		a = append(a, rand.Intn(int(n))*(2*rand.Intn(2)-1))
	}
	os.Exit(m.Run())
}

func TestSubarrayDC(t *testing.T) {
	l, h := MaxSumBruteForce(a)
	ll, hh := MaxSumDC(a)
	if l != ll && h != hh {
		t.Errorf("Brute-Force: (%d, %d), DC: (%d, %d)", l, h, ll, hh)
	}
}

func BenchmarkSubarrayBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSumBruteForce(a)
	}
}
