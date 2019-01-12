package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	arr := copyArr()
	SelectionSort(arr)
	testSorted(t, arr)
	testContainsAll(t, arr)
}

func BenchmarkSelectionSort(b *testing.B) {
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectionSort(arr)
	}
}
