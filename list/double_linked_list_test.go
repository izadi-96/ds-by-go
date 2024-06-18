package list_test

import (
	"ds-by-go/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertBeginning(t *testing.T) {
	sample := &list.DoubleLinkedList{}
	sample.InsertBeginning(1)
	sample.InsertBeginning(2)
	sample.InsertBeginning(3)

	result, _ := sample.GetAtPosition(1)
	assert.Equal(t, 3, result)
	assert.Equal(t, 3, sample.Length())
}

func TestInsertEnd(t *testing.T) {
	sample := &list.DoubleLinkedList{}
	sample.InsertEnd(1)
	sample.InsertEnd(2)
	sample.InsertEnd(3)

	result, _ := sample.GetAtPosition(1)
	assert.Equal(t, 1, result)
	assert.Equal(t, 3, sample.Length())
}

func TestInsert(t *testing.T) {
	sample := &list.DoubleLinkedList{}
	sample.InsertEnd(1)
	sample.InsertEnd(2)
	sample.InsertEnd(3)
	sample.Insert(2, 8)

	sample.Display()
}
