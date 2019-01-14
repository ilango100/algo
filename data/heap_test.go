package data

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := Heap([]int{1, 2, 3, 4, 5, 6, 7})
	h.BuildMax()
	fmt.Println(h)
}

func TestMinHeap(t *testing.T) {
	h := Heap([]int{7, 6, 5, 4, 3, 2, 1})
	h.BuildMin()
	fmt.Println(h)
}
