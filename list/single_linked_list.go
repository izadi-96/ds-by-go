package list

import (
	"fmt"
)

type SingleLinkedNode struct {
	Date interface{}
	Next *SingleLinkedNode
}

type SingleLinkedList struct {
	head *SingleLinkedNode
	size int
}

func (ll *SingleLinkedList) Display() error {
	if ll.head == nil {
		return fmt.Errorf("display: List is empty")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.Date)
		current = current.Next
	}
	fmt.Println()
	return nil
}

func (ll *SingleLinkedList) Length() int {
	return ll.size
}

func (ll *SingleLinkedList) GetAtPosition(position int) (interface{}, error) {
	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("insert: Index out of bounds")
	}

	current := ll.head
	for position > 1 {
		current = current.Next
	}
	return current.Date, nil
}

func (ll *SingleLinkedList) InsertBeginning(data interface{}) {
	node := &SingleLinkedNode{
		Date: data,
		Next: nil,
	}

	if ll.head == nil {
		ll.head = node
	} else {
		node.Next = ll.head
		ll.head = node
	}
	ll.size++
}

func (ll *SingleLinkedList) InsertEnd(data interface{}) {
	node := &SingleLinkedNode{
		Date: data,
		Next: nil,
	}
	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	ll.size++
}

func (ll *SingleLinkedList) Insert(position int, data interface{}) error {
	if position < 1 || position > ll.size+1 {
		return fmt.Errorf("insert: Index out of bounds")
	}
	node := &SingleLinkedNode{
		Date: data,
		Next: nil,
	}
	var prev, current *SingleLinkedNode
	current = ll.head
	prev = nil

	for position > 1 {
		prev = current
		current = current.Next
		position = position - 1
	}

	if prev == nil {
		node.Next = current
		ll.head = node
	} else {
		node.Next = current
		prev.Next = node
	}
	ll.size++
	return nil
}

func (ll *SingleLinkedList) DeleteFirst() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}
	current := ll.head
	ll.head = ll.head.Next
	ll.size--
	return current.Date, nil
}

func (ll *SingleLinkedList) DeleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}

	var current, prev *SingleLinkedNode
	current = ll.head
	prev = nil
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	if prev == nil {
		ll.head = nil
	} else {
		prev.Next = nil
	}
	ll.size--
	return current.Date, nil
}

func (ll *SingleLinkedList) DeleteAt(position int) (interface{}, error) {
	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("insert: Index out of bounds")
	}

	var current, prev *SingleLinkedNode
	current = ll.head
	prev = nil
	for position > 1 {
		prev = current
		current = current.Next
		position--
	}

	if prev == nil {
		ll.head = ll.head.Next
	} else {
		prev.Next = current.Next
	}
	ll.size--
	return current.Date, nil
}

func (ll *SingleLinkedList) FindKthFromEnd(n int) interface{} {
	first, second := ll.head, ll.head
	for ; n > 0; n-- {
		second = second.Next
	}
	for ; second.Next != nil; first, second = first.Next, second.Next {
	}
	return first.Next.Date
}

func (ll *SingleLinkedList) Link(node1 *SingleLinkedNode, node2 *SingleLinkedNode) {
	if ll.head == nil {
		ll.head = node1
	}
	node1.Next = node2
	ll.size++
}
