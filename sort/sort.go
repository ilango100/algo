package sort

//Sorter interface provides sorting function
type Sorter interface {
	Sort([]int)
}

//Func constructs Sorter from a function
type Func func([]int)

//Sort implements Sorter interface
func (s Func) Sort(a []int) {
	s(a)
}
