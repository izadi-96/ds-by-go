package list

import (
	"testing"
)

func TestCircularLinkedListInsertBeginning(t *testing.T) {
	sample := &CircularLinkedList{}
	sample.InsertBeginning(1)
	sample.InsertBeginning(3)
	sample.InsertBeginning(5)
	sample.Display()
}

func TestCircularLinkedListInsertEnd(t *testing.T) {
	sample := &CircularLinkedList{}
	sample.InsertBeginning(1)
	sample.InsertBeginning(3)
	sample.InsertBeginning(5)
	sample.InsertEnd(10)
	sample.Display()
}

func TestCircularLinkedListInsertAt(t *testing.T) {
	sample := &CircularLinkedList{}
	sample.InsertBeginning(1)
	sample.InsertBeginning(3)
	sample.InsertBeginning(5)
	sample.Insert(2, 10)
	sample.Display()
}

func TestCircularLinkedListDeleteEnd(t *testing.T) {
	sample := &CircularLinkedList{}
	sample.InsertBeginning(1)
	sample.InsertBeginning(3)
	sample.InsertBeginning(5)
	sample.DeleteLast()
	sample.Display()
}
