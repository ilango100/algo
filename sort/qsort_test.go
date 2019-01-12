package sort

import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := copyArr()
	QuickSort(arr, PivotMedian)
	testSorted(t, arr)
	testContainsAll(t, arr)
}

func BenchmarkQuickSort(b *testing.B) {
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr, PivotMedian)
	}
}

func BenchmarkQuickSortFirst(b *testing.B) {
	if !testing.Verbose() {
		b.SkipNow()
	}
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr, PivotFirst)
	}
}

func BenchmarkQuickSortLast(b *testing.B) {
	if !testing.Verbose() {
		b.SkipNow()
	}
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr, PivotLast)
	}
}

func BenchmarkGoSort(b *testing.B) {
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(arr)
	}
}
