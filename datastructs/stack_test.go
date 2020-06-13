package datastructs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := Stack{}

	s.Push(1)

	s.Push(2)

	n, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, n)

	s.Push(3)

	n, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 3, n)

	n, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, n)
}
