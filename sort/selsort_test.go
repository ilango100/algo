package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	testFunc(t, SelectionSort)
}

func BenchmarkSelectionSort(b *testing.B) {
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectionSort(arr)
	}
}
