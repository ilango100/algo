package algo

import (
	"flag"
	"math/rand"
	"os"
	"sort"
	"testing"
)

var a []int

func TestMain(m *testing.M) {
	n := flag.Uint("n", 1000, "Number of inputs")
	flag.Parse()
	a = rand.Perm(int(*n))
	os.Exit(m.Run())
}

func TestQSort(t *testing.T) {
	arr := make([]int, len(a))
	copy(arr, a)
	QSort(arr, MedianPiv)
	e := len(arr) - 2
	for i := 0; i < e; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("%d [%d] came before %d [%d]:", arr[i], i, arr[i+1], i+1)
		}
	}
}

func BenchmarkQSort(b *testing.B) {
	arr := make([]int, len(a))
	copy(arr, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QSort(arr, MedianPiv)
	}
}

func BenchmarkQSortFirst(b *testing.B) {
	if !testing.Verbose() {
		b.SkipNow()
	}
	arr := make([]int, len(a))
	copy(arr, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QSort(arr, FirstPiv)
	}
}

func BenchmarkQSortLast(b *testing.B) {
	if !testing.Verbose() {
		b.SkipNow()
	}
	arr := make([]int, len(a))
	copy(arr, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QSort(arr, LastPiv)
	}
}

func BenchmarkGoSort(b *testing.B) {
	arr := make([]int, len(a))
	copy(arr, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(arr)
	}
}
