package datastructs

import "errors"

// Stack represents a stack of integers
// The first element is the "bottom" of the stack, the last element is the "top"
type Stack []int

// Push adds the given value to the top of the stack
func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

// Pop takes out the top value from the stack.
// An error is returned if there are no values in the stack
func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("no elements to pop")
	}

	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n, nil
}
