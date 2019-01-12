package sort

//InsertionSort sorts using the Insertion sort: https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort(a []int) {
	for i := range a {
		for j := 0; j < i; j++ {
			if a[j] > a[i] {
				tmp := a[i]
				copy(a[j+1:i+1], a[j:i])
				a[j] = tmp
				break
			}
		}
	}
}
