package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDynamicStackResize(t *testing.T) {
	sample := NewDynamicStack(4)
	sample.Push(1)
	sample.Push(2)
	sample.Push(4)
	sample.Pop()
	sample.Pop()

	assert.Equal(t, uint(2), sample.Capacity())
}
