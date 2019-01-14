package data

/*A Queue is a FIFO collection.
You can use a prefilled []int to start with those elements.
	Queue([]int{1,2,3}).Dequeue() //=>1*/
type Queue []int

/*Enqueue adds a new element in the queue.
Returns the same Queue for chaining, i.e: q.Enqueue(1).Enqueue(2) */
func (q *Queue) Enqueue(i int) *Queue {
	*q = append(*q, i)
	return q
}

/*Dequeue extracts the next element from the queue*/
func (q *Queue) Dequeue() int {
	r := (*q)[0]
	*q = (*q)[1:]
	return r
}
