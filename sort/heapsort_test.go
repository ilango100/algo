package sort

import "testing"

func TestHeapSort(t *testing.T) {
	testFunc(t, HeapSort)
}

func BenchmarkSortHeap(b *testing.B) {
	benchFunc(b, HeapSort)
}
