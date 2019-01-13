package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	testFunc(t, SelectionSort)
}

func BenchmarkSortSelection(b *testing.B) {
	benchFunc(b, SelectionSort)
}
