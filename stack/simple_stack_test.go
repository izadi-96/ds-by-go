package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushSimpleStack(t *testing.T) {
	fmt.Println("salam 1")
	sample := NewSimpleStack(10)
	sample.Push(1)
	sample.Push(2)

	result, _ := sample.Peek()
	assert.Equal(t, 2, result)
}

func TestIsFullSimpleStack(t *testing.T) {
	fmt.Println("salam 2")
	sample := NewSimpleStack(3)
	sample.Push(1)
	sample.Push(2)
	sample.Push(3)

	err := sample.Push(2)
	assert.Errorf(t, err, "stack is full")
}

func TestIsEmptySimpleStack(t *testing.T) {
	sample := NewSimpleStack(3)

	_, err := sample.Pop()
	assert.Errorf(t, err, "stack is empty")
}
