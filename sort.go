package algo

//Sorter interface provides sorting function
type Sorter interface {
	Sort([]int)
}

//SortFunc constructs Sorter from a function
type SortFunc func([]int)

//Sort implements Sorter interface
func (s SortFunc) Sort(a []int) {
	s(a)
}
