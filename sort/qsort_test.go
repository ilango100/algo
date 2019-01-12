package sort

import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := copyArr()
	QuickSort(arr, MedianPiv)
	testSorted(t, arr)
	testContainsAll(t, arr)
}

func BenchmarkQuickSort(b *testing.B) {
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr, MedianPiv)
	}
}

func BenchmarkQuickSortFirst(b *testing.B) {
	if !testing.Verbose() {
		b.SkipNow()
	}
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr, FirstPiv)
	}
}

func BenchmarkQuickSortLast(b *testing.B) {
	if !testing.Verbose() {
		b.SkipNow()
	}
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr, LastPiv)
	}
}

func BenchmarkGoSort(b *testing.B) {
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(arr)
	}
}
