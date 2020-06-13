package datastructs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := Queue{}

	q.Enqueue(1)

	q.Enqueue(2)

	n, err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 1, n)

	q.Enqueue(3)

	n, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 2, n)

	n, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 3, n)
}
