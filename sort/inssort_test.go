package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	testFunc(t, InsertionSort)
}

func BenchmarkInsertionSort(b *testing.B) {
	benchFunc(b, InsertionSort)
}
