package sort

/*CountingSort implementation https://en.wikipedia.org/wiki/Counting_sort.
Counting Sort is non-comparison sort. The []int should have values > 0.*/
func CountingSort(b []int) {
	a := make([]int, len(b))
	copy(a, b)
	k := 0
	for _, v := range a {
		if v > k {
			k = v
		}
	}
	c := make([]int, k+1)
	for i := range a {
		c[a[i]]++
	}
	for i := range c {
		if i == 0 {
			continue
		}
		c[i] = c[i] + c[i-1]
	}
	for i := len(a) - 1; i >= 0; i-- {
		c[a[i]]--
		b[c[a[i]]] = a[i]
	}
}
