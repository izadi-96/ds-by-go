package list

import (
	"fmt"
)

type DoubleLinkNode struct {
	data interface{}
	next *DoubleLinkNode
	prev *DoubleLinkNode
}

func (n *DoubleLinkNode) Get() interface{} {
	return n
}

type DoubleLinkedList struct {
	head *DoubleLinkNode
	tail *DoubleLinkNode
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

//func (ll *DoubleLinkedList) checkIfEmptyAndAdd(newNode *DoubleLinkNode) bool {
//	if ll.size == 0 {
//		// insert first node in doubly linked list
//		ll.head = newNode
//		ll.tail = newNode
//		ll.size++
//		return true
//	}
//	return false
//}

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
	node := &DoubleLinkNode{
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
	node := &DoubleLinkNode{
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

	node := &DoubleLinkNode{
		data: data,
	}
	node.prev = current.prev
	node.next = current

	current.prev.next = node
	current.next.prev = node
	return nil
}

func (ll *DoubleLinkedList) DeleteFirst() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}

	data := ll.head.data
	if ll.size == 1 {
		ll.head = nil
		ll.tail = nil
		ll.size = 0
		return data, nil
	}

	ll.head.next.prev = nil
	ll.head = ll.head.next
	ll.size--
	return data, nil
}

func (ll *DoubleLinkedList) DeleteLast() (interface{}, error) {
	if ll.tail == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}

	current := ll.tail
	if ll.size == 1 {
		ll.head = nil
		ll.tail = nil
		ll.size = 0
		return current.data, nil
	}

	ll.tail.prev.next = nil
	ll.tail = ll.tail.prev
	ll.size--
	return current.data, nil
}

func (ll *DoubleLinkedList) DeleteAt(position int) (interface{}, error) {
	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("insert: Index out of bounds")
	}
	var current, prev *DoubleLinkNode
	current = ll.head
	prev = nil

	for position > 1 {
		prev = current
		current = current.next
		position--
	}

	prev.next = current.next
	current.next.prev = prev
	ll.size--
	return current.data, nil
}
