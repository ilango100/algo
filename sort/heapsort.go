package sort

import (
	"github.com/ilango100/algo/data"
)

/*HeapSort sorts by Heap Sort: https://en.wikipedia.org/wiki/Heapsort.*/
func HeapSort(a []int) {
	h := data.Heap(a)
	h.BuildMax()
	for i := len(a) - 1; i >= 0; i-- {
		h.Swap(0, i)
		h = data.Heap(a[:i])
		h.MaxHeapify(0)
	}
}
