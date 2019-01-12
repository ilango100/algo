package sort

//QuickSort is Quicksort implementation: https://en.wikipedia.org/wiki/Quicksort
//
//The pivot function should return the index of pivot and not the value!
func QuickSort(a []int, pivot func([]int) int) {
	if len(a) <= 1 {
		return
	}
	p := pivot(a)
	for i := 0; i < len(a); i++ {
		if a[p] < a[i] && p > i {
			a[i], a[p-1], a[p] = a[p-1], a[p], a[i]
			p--
			i--
		} else if a[p] > a[i] && p < i {
			a[i], a[p+1], a[p] = a[p+1], a[p], a[i]
			p++
			i--
		}
	}
	QuickSort(a[:p], pivot)
	QuickSort(a[p+1:], pivot)
}

//PivotLast always returns last element as pivot, i.e., len(a)-1
func PivotLast(a []int) int {
	return len(a) - 1
}

//PivotFirst always returns first element as pivot, i.e., 0
func PivotFirst(a []int) int {
	return 0
}

//PivotMedian takes the median of first, middle and last element of the slice
func PivotMedian(a []int) int {
	m := len(a) / 2
	l := len(a) - 1
	b := a[0]
	c := a[m]
	d := a[l]
	if b < c {
		if c < d {
			return m
		} else if d < b {
			return 0
		} else {
			return l
		}
	} else {
		if c > d {
			return m
		} else if d > b {
			return 0
		} else {
			return l
		}
	}
}

//QuickSorter is the quicksort implementation of Sorter interface.
//It is of same type as its pivot function:
//	QuickSorter(MedianPiv).Sort([]int{1,2,3})
//Don't use inline slices, since you cannot retrieve it back.
type QuickSorter func([]int) int

//Sort is implementation of Sorter interface
func (q QuickSorter) Sort(a []int) {
	QuickSort(a, q)
}
