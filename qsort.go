package algo

//QSort is Quicksort implementation: https://en.wikipedia.org/wiki/Quicksort
//
//The pivot function should return the index of pivot and not the value!
func QSort(a []int, pivot func([]int) int) {
	if len(a) <= 1 {
		return
	}
	p := pivot(a)
	for i := 0; i < len(a); i++ {
		if a[p] < a[i] && p > i {
			tmp := a[i]
			a[i] = a[p-1]
			a[p-1] = a[p]
			a[p] = tmp
			p--
			i--
		} else if a[p] > a[i] && p < i {
			tmp := a[i]
			a[i] = a[p+1]
			a[p+1] = a[p]
			a[p] = tmp
			p++
			i--
		}
	}
	QSort(a[:p], pivot)
	QSort(a[p+1:], pivot)
}

//LastPiv always returns last element as pivot, i.e., len(a)-1
func LastPiv(a []int) int {
	return len(a) - 1
}

//FirstPiv always returns first element as pivot, i.e., 0
func FirstPiv(a []int) int {
	return 0
}

//MedianPiv takes the median of first, middle and last element of the slice
func MedianPiv(a []int) int {
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

//QSorter is the quicksort implementation of Sorter interface.
//It is of same type as its pivot function:
//	QSorter(MedianPiv).Sort([]int{1,2,3})
//
type QSorter func([]int) int

//Sort is implementation of Sorter interface
func (q QSorter) Sort(a []int) {
	QSort(a, q)
}
