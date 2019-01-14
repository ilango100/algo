package data

/*A Heap is implementation of heap data structure: https://en.wikipedia.org/wiki/Binary_heap
This is a generic heap. Call BuildMax or BuildMin to build Max-Heap or Min-Heap.
Heap provides direct access to all its elements. Any operation: Heapify or Build changes the order of elements. */
type Heap []int

func (h Heap) left(i int) int {
	return i<<1 + 1
}

func (h Heap) right(i int) int {
	return i<<1 + 2
}

/*Swap swaps values between two indices*/
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

/*MaxHeapify "heapifies" the heap at given index.
https://en.wikipedia.org/wiki/Binary_heap#Heap_operations
This is for internal use, but exported for study and research. */
func (h Heap) MaxHeapify(i int) {
	if i >= len(h)/2 {
		return
	}
	l, r := h.left(i), h.right(i)
	if h[l] > h[i] && h[l] > h[r] {
		h.Swap(i, l)
		h.MaxHeapify(l)
	} else if h[r] > h[i] && h[r] > h[l] {
		h.Swap(i, r)
		h.MaxHeapify(r)
	}
}

/*BuildMax builds a Max-Heap from the given slice.*/
func (h Heap) BuildMax() {
	for i := len(h)/2 - 1; i >= 0; i-- {
		h.MaxHeapify(i)
	}
}

/*MinHeapify "heapifies" the heap at given index.
https://en.wikipedia.org/wiki/Binary_heap#Heap_operations
This is for internal use, but exported for study and research. */
func (h Heap) MinHeapify(i int) {
	if i >= len(h)/2 {
		return
	}
	l, r := h.left(i), h.right(i)
	if h[l] < h[i] && h[l] < h[r] {
		h.Swap(i, l)
		h.MinHeapify(l)
	} else if h[r] < h[i] && h[r] < h[l] {
		h.Swap(i, r)
		h.MinHeapify(r)
	}
}

/*BuildMin builds a Min-Heap from the given slice.*/
func (h Heap) BuildMin() {
	for i := len(h)/2 - 1; i >= 0; i-- {
		h.MinHeapify(i)
	}
}
