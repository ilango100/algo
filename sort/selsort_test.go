package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	arr := copyArr()
	SelectionSort(arr)
	testSorted(t, arr)
	testContainsAll(t, arr)
}
