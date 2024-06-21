package list

import (
	"fmt"
)

type SingleLinkedNode struct {
	data interface{}
	next *SingleLinkedNode
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
		fmt.Printf("%v -> ", current.data)
		current = current.next
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
		current = current.next
	}
	return current.data, nil
}

func (ll *SingleLinkedList) InsertBeginning(data interface{}) {
	node := &SingleLinkedNode{
		data: data,
		next: nil,
	}

	if ll.head == nil {
		ll.head = node
	} else {
		node.next = ll.head
		ll.head = node
	}
	ll.size++
}

func (ll *SingleLinkedList) InsertEnd(data interface{}) {
	node := &SingleLinkedNode{
		data: data,
		next: nil,
	}
	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	ll.size++
}

func (ll *SingleLinkedList) Insert(position int, data interface{}) error {
	if position < 1 || position > ll.size+1 {
		return fmt.Errorf("insert: Index out of bounds")
	}
	node := &SingleLinkedNode{
		data: data,
		next: nil,
	}
	var prev, current *SingleLinkedNode
	current = ll.head
	prev = nil

	for position > 1 {
		prev = current
		current = current.next
		position = position - 1
	}

	if prev == nil {
		node.next = current
		ll.head = node
	} else {
		node.next = current
		prev.next = node
	}
	ll.size++
	return nil
}

func (ll *SingleLinkedList) DeleteFirst() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}
	current := ll.head
	ll.head = ll.head.next
	ll.size--
	return current.data, nil
}

func (ll *SingleLinkedList) DeleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}

	var current, prev *SingleLinkedNode
	current = ll.head
	prev = nil
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev == nil {
		ll.head = nil
	} else {
		prev.next = nil
	}
	ll.size--
	return current.data, nil
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
		current = current.next
		position--
	}

	if prev == nil {
		ll.head = ll.head.next
	} else {
		prev.next = current.next
	}
	ll.size--
	return current.data, nil
}

func (ll *SingleLinkedList) FindKthFromEnd(n int) interface{} {
	first, second := ll.head, ll.head
	for ; n > 0; n-- {
		second = second.next
	}
	for ; second.next != nil; first, second = first.next, second.next {
	}
	return first.next.data
}
