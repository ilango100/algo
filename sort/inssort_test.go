package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	testFunc(t, InsertionSort)
}

func BenchmarkSortInsertion(b *testing.B) {
	benchFunc(b, InsertionSort)
}
