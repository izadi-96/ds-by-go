package list

import (
	"fmt"
)

type doubleLinkNode struct {
	data interface{}
	next *doubleLinkNode
	prev *doubleLinkNode
}

type DoubleLinkedList struct {
	head *doubleLinkNode
	tail *doubleLinkNode
	size int
}

func (ll *DoubleLinkedList) Display() error {
	if ll.head == nil {
		return fmt.Errorf("display: List is empty")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}

func (ll *DoubleLinkedList) Length() int {
	return ll.size
}

func (ll *DoubleLinkedList) GetHead() interface{} {
	return ll.head.data
}

func (ll *DoubleLinkedList) GetTail() interface{} {
	return ll.tail.data
}

func (ll *DoubleLinkedList) checkIfEmptyAndAdd(newNode *doubleLinkNode) bool {
	if ll.size == 0 {
		// insert first node in doubly linked list
		ll.head = newNode
		ll.tail = newNode
		ll.size++
		return true
	}
	return false
}

func (ll *DoubleLinkedList) GetAtPosition(position int) (interface{}, error) {
	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("insert: Index out of bounds")
	}

	current := ll.head
	for position > 1 {
		current = current.next
		position = position - 1
	}
	return current.data, nil
}

func (ll *DoubleLinkedList) InsertBeginning(data interface{}) {
	node := &doubleLinkNode{
		data: data,
		next: nil,
		prev: nil,
	}
	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {

		ll.Display()
		node.next = ll.head
		ll.head.prev = node
		ll.head = node
	}
	ll.size++
}

func (ll *DoubleLinkedList) InsertEnd(data interface{}) {
	node := &doubleLinkNode{
		data: data,
		next: nil,
		prev: nil,
	}

	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		node.prev = ll.tail
		ll.tail = node
	}
	ll.size++
}

func (ll *DoubleLinkedList) Insert(position int, data interface{}) error {

	if position == 1 || ll.size == 0 {
		ll.InsertBeginning(data)
		return nil
	}

	if position == ll.size {
		ll.InsertEnd(data)
		return nil
	}

	current := ll.head
	for position > 1 {
		current = current.next
		position--
	}

	node := &doubleLinkNode{
		data: data,
	}
	node.prev = current.prev
	node.next = current

	current.prev.next = node
	current.next.prev = node
	return nil
}
