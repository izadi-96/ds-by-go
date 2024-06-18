package list_test

import (
	"ds-by-go/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertBeginningDLL(t *testing.T) {
	sample := &list.DoubleLinkedList{}
	sample.InsertBeginning(1)
	sample.InsertBeginning(2)
	sample.InsertBeginning(3)

	result, _ := sample.GetAtPosition(1)
	assert.Equal(t, 3, result)
	assert.Equal(t, 3, sample.Length())
}

func TestInsertEndDLL(t *testing.T) {
	sample := &list.DoubleLinkedList{}
	sample.InsertEnd(1)
	sample.InsertEnd(2)
	sample.InsertEnd(3)

	result, _ := sample.GetAtPosition(1)
	assert.Equal(t, 1, result)
	assert.Equal(t, 3, sample.Length())
}

func TestInsertDLL(t *testing.T) {
	sample := &list.DoubleLinkedList{}
	sample.InsertEnd(1)
	sample.InsertEnd(2)
	sample.InsertEnd(3)
	sample.Insert(2, 8)

	sample.Display()
}

func TestDeleteFirstDLL(t *testing.T) {
	sample := &list.DoubleLinkedList{}
	sample.InsertEnd(1)
	sample.InsertEnd(2)
	sample.InsertEnd(3)

	result, _ := sample.DeleteFirst()
	sample.Display()
	assert.Equal(t, 1, result)
}

func TestDeleteLastDLL(t *testing.T) {
	sample := &list.DoubleLinkedList{}
	sample.InsertEnd(1)
	sample.InsertEnd(2)
	sample.InsertEnd(3)

	result, _ := sample.DeleteLast()
	sample.Display()
	assert.Equal(t, 3, result)
}

func TestDeleteAtDLL(t *testing.T) {
	sample := &list.DoubleLinkedList{}
	sample.InsertEnd(1)
	sample.InsertEnd(2)
	sample.InsertEnd(3)
	sample.InsertEnd(4)
	sample.InsertEnd(5)

	result, _ := sample.DeleteAt(3)
	sample.Display()
	assert.Equal(t, 3, result)
}
