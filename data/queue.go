package data

/*A Queue is a FIFO collection.*/
type Queue struct {
	arr []int
}

/*Enqueue adds ("push") a new element in the queue.
Returns the same Queue for chaining, i.e: q.Enqueue(1).Enqueue(2) */
func (q *Queue) Enqueue(i int) *Queue {
	q.arr = append(q.arr, i)
	return q
}

/*Dequeue extracts the next element from the queue*/
func (q *Queue) Dequeue() int {
	r := q.arr[0]
	q.arr = q.arr[1:]
	return r
}
