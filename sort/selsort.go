package sort

//SelectionSort sorts using Selection Sort: https://en.wikipedia.org/wiki/Selection_sort
func SelectionSort(a []int) {
	min := a[0]
	mini := 0
	for i := range a {
		min = a[i]
		mini = i
		for j := i; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
				mini = j
			}
		}
		a[mini] = a[i]
		a[i] = min
	}
}
