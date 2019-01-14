package data

/*A Stack is a LIFO collection.
You can use a prefilled []int to start with those elements.
	Stack([]int{1,2,3}).Pop() //=>3*/
type Stack []int

/*Push pushes an element into the stack.
Returns stack for chaining, i.e: s.Push(1).Push(1).*/
func (s *Stack) Push(i int) *Stack {
	*s = append(*s, i)
	return s
}

/*Pop pops the next element from the stack.*/
func (s *Stack) Pop() int {
	l := len(*s) - 1
	r := (*s)[l]
	*s = (*s)[:l]
	return r
}
