package algo

import (
	"math/rand"
	"sort"
	"testing"
)

var a = rand.Perm(1000)

func TestQSort(t *testing.T) {
	arr := make([]int, len(a))
	copy(arr, a)
	QSort(arr, FirstPiv)
	e := len(arr) - 2
	for i := 0; i < e; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("%d [%d] came before %d [%d]:", arr[i], i, arr[i+1], i+1)
		}
	}
}

func BenchmarkQSortMed(b *testing.B) {
	arr := make([]int, len(a))
	copy(arr, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QSort(arr, MedianPiv)
	}
}

func BenchmarkQSortFirst(b *testing.B) {
	arr := make([]int, len(a))
	copy(arr, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QSort(arr, FirstPiv)
	}
}

func BenchmarkQSortLast(b *testing.B) {
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
