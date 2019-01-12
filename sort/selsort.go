package sort

//SelectionSort sorts using Selection Sort: https://en.wikipedia.org/wiki/Selection_sort
//
//You can wrap SelectionSort with Func to get Sorter:
//	Func(SelectionSort).Sort([]int{3,2,1})
//Don't use inline slices, since you cannot retrieve it back.
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
