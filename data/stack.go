package data

/*A Stack is a LIFO collection.*/
type Stack struct {
	arr []int
}

/*Push pushes an element into the stack.
Returns stack for chaining, i.e: s.Push(1).Push(1).*/
func (s *Stack) Push(i int) *Stack {
	s.arr = append(s.arr, i)
	return s
}

/*Pop pops the next element from the stack.*/
func (s *Stack) Pop() int {
	l := len(s.arr) - 1
	r := s.arr[l]
	s.arr = s.arr[:l]
	return r
}
