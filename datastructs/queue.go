package datastructs

import "errors"

// Queue represents a queue of integers
type Queue []int

// Enqueue adds the given value to the end of the queue
func (q *Queue) Enqueue(n int) {
	*q = append(*q, n)
}

// Dequeue takes out the first value from the queue.
// An error is returned if there are no values in the queue
func (q *Queue) Dequeue() (int, error) {
	if len(*q) == 0 {
		return 0, errors.New("no elements to dequeue")
	}

	n := (*q)[0]
	*q = (*q)[1:]
	return n, nil
}
