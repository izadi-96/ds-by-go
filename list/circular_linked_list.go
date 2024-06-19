package list

import (
	"fmt"
)

type circularLinkNode struct {
	data interface{}
	next *circularLinkNode
}

type CircularLinkedList struct {
	size int
	head *circularLinkNode
}

func (ll *CircularLinkedList) Length() int {
	if ll.head == nil {
		return 0
	}
	current := ll.head.next
	count := 1
	for current != ll.head {
		current = current.next
		count++
	}
	return count
}

func (ll *CircularLinkedList) Display() {
	head := ll.head
	for i := 0; i < ll.size; i++ {
		fmt.Print(head.data)
		fmt.Print("-->")
		head = head.next
	}
	fmt.Println()
}

func (ll *CircularLinkedList) InsertBeginning(data interface{}) {
	node := &circularLinkNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = node
		ll.head.next = node
		ll.size++
	} else {
		current := ll.head
		for {
			if current.next == ll.head {
				break
			}
			current = current.next
		}
		current.next = node
		node.next = ll.head
		ll.head = node
		ll.size++
	}
}

func (ll *CircularLinkedList) InsertEnd(data interface{}) {
	node := &circularLinkNode{
		data: data,
	}

	if ll.head == nil {
		ll.head = node
		ll.head.next = node
		ll.size++
	} else {
		current := ll.head
		for {
			if current.next == ll.head {
				break
			}
			current = current.next
		}
		current.next = node
		node.next = ll.head
		ll.size++
	}
}

func (ll *CircularLinkedList) Insert(position int, data interface{}) error {
	if position == 1 {
		ll.InsertBeginning(data)
		return nil
	}

	if position == ll.size {
		ll.InsertEnd(data)
		return nil
	}

	node := &circularLinkNode{
		data: data,
	}
	var current, prev *circularLinkNode
	current = ll.head
	prev = nil
	for position > 1 {
		prev = current
		current = current.next
		position--
	}
	prev.next = node
	node.next = current
	ll.size++
	return nil
}

func (ll *CircularLinkedList) DeleteFirst() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}
	current := ll.head
	ll.head = ll.head.next
	ll.size--
	return current.data, nil
}

func (ll *CircularLinkedList) DeleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}
	var current, prev *circularLinkNode
	current = ll.head
	prev = nil
	for {
		prev = current
		if current.next == ll.head {
			break
		}
		current = current.next
	}
	prev.next = nil
	ll.size--
	return current.data, nil
}
