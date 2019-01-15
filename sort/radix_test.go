package sort

import (
	"testing"
)

func TestSortRadix(t *testing.T) {
	testFunc(t, RadixSort)
}

func BenchmarkSortRadix(b *testing.B) {
	benchFunc(b, RadixSort)
}
