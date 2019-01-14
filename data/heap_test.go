package data

import (
	"math/rand"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := Heap(rand.Perm(1000))
	h.BuildMax()
}

func TestMinHeap(t *testing.T) {
	h := Heap(rand.Perm(1000))
	h.BuildMin()
}
