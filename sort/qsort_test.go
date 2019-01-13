package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := copyArr()
	QuickSort(arr, PivotMedian)
	testSorted(t, arr)
	testContainsAll(t, arr)
}

func TestQuickSortFirst(t *testing.T) {
	arr := copyArr()
	QuickSort(arr, PivotFirst)
	testSorted(t, arr)
	testContainsAll(t, arr)
}

func TestQuickSortLast(t *testing.T) {
	arr := copyArr()
	QuickSort(arr, PivotLast)
	testSorted(t, arr)
	testContainsAll(t, arr)
}

func BenchmarkSortQuick(b *testing.B) {
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr, PivotMedian)
	}
}

func BenchmarkSortQuickFirst(b *testing.B) {
	if !testing.Verbose() {
		b.SkipNow()
	}
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr, PivotFirst)
	}
}

func BenchmarkSortQuickLast(b *testing.B) {
	if !testing.Verbose() {
		b.SkipNow()
	}
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr, PivotLast)
	}
}
