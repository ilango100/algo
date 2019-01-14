package data

/*A Heap is implementation of heap data structure: https://en.wikipedia.org/wiki/Binary_heap
This is a generic heap. Call BuildMax or BuildMin to build Max-Heap or Min-Heap. */
type Heap []int

func (h Heap) left(i int) int {
	return (i+1)<<1 - 1
}

func (h Heap) right(i int) int {
	return (i + 1) << 1
}

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) maxheapify(i int) {
	if i >= len(h)/2 {
		return
	}
	l, r := h.left(i), h.right(i)
	if h[l] > h[i] && h[l] > h[r] {
		h.swap(i, l)
		h.maxheapify(l)
	} else if h[r] > h[i] && h[r] > h[l] {
		h.swap(i, r)
		h.maxheapify(r)
	}
}

/*BuildMax builds a Max-Heap from the given slice.*/
func (h Heap) BuildMax() {
	for i := len(h)/2 - 1; i >= 0; i-- {
		h.maxheapify(i)
	}
}

func (h Heap) minheapify(i int) {
	if i >= len(h)/2 {
		return
	}
	l, r := h.left(i), h.right(i)
	if h[l] < h[i] && h[l] < h[r] {
		h.swap(i, l)
		h.minheapify(l)
	} else if h[r] < h[i] && h[r] < h[l] {
		h.swap(i, r)
		h.minheapify(r)
	}
}

/*BuildMin builds a Min-Heap from the given slice.*/
func (h Heap) BuildMin() {
	for i := len(h)/2 - 1; i >= 0; i-- {
		h.minheapify(i)
	}
}
