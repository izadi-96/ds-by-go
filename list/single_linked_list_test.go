package list_test

import (
	"ds-by-go/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createSample() *list.SingleLinkedList {
	ll := &list.SingleLinkedList{}
	ll.InsertEnd(1)
	ll.InsertEnd(2)
	ll.InsertEnd(3)
	ll.InsertEnd(4)
	ll.InsertEnd(5)
	ll.InsertEnd(6)
	return ll
}

func TestDeleteAt(t *testing.T) {
	sample := createSample()
	result, _ := sample.DeleteAt(3)

	assert.Equal(t, 3, result)
}

func TestDeleteFirst(t *testing.T) {
	sample := createSample()
	result, _ := sample.DeleteFirst()

	assert.Equal(t, 1, result)
}

func TestDeleteLast(t *testing.T) {
	sample := createSample()
	result, _ := sample.DeleteLast()

	assert.Equal(t, 6, result)
}
