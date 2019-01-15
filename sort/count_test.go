package sort

import "testing"

func TestCountingSort(t *testing.T) {
	testFunc(t, CountingSort)
}

func BenchmarkSortCounting(b *testing.B) {
	benchFunc(b, CountingSort)
}
