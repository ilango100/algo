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
	s, l, h := MaxSumBruteForce(a)
	ss, ll, hh := MaxSumDC(a)
	//Check only sum since there can be multiple subarrays
	if s != ss {
		t.Errorf("Brute-Force: (%d, %d, sum: %d) , DC: (%d, %d, sum: %d)", l, h, s, ll, hh, ss)
	}
}

func TestSubarrayKadane(t *testing.T) {
	s, l, h := MaxSumBruteForce(a)
	ss, ll, hh := MaxSumKadane(a)
	//Check only sum since there can be multiple subarrays
	if s != ss {
		t.Errorf("Brute-Force: (%d, %d, sum: %d) , Kadane: (%d, %d, sum: %d)", l, h, s, ll, hh, ss)
	}
}

func BenchmarkSubarrayBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSumBruteForce(a)
	}
}

func BenchmarkSubarrayDC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSumDC(a)
	}
}

func BenchmarkSubarrayKadane(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSumKadane(a)
	}
}
