package data

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := Heap([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	h.BuildMax()
	fmt.Println(h)
}

func TestMinHeap(t *testing.T) {
	h := Heap([]int{7, 6, 5, 4, 3, 2, 1, 8, 9, 10, 11, 12})
	h.BuildMin()
	fmt.Println(h)
}
