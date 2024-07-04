package stack

import "errors"

type DynamicStack[T any] struct {
	top      int
	capacity uint
	data     []interface{}
}

func (s *DynamicStack[T]) Init(capacity uint) *DynamicStack[T] {
	s.top = -1
	s.capacity = capacity
	s.data = make([]interface{}, capacity)
	return s
}

func NewDynamicStack[T any](capacity uint) *DynamicStack[T] {
	return new(DynamicStack[T]).Init(capacity)
}

func (s *DynamicStack[T]) IsEmpty() bool {
	return s.top == -1
}

func (s *DynamicStack[T]) IsFull() bool {
	return s.top == int(s.capacity)-1
}

func (s *DynamicStack[T]) Size() uint {
	return uint(s.top + 1)
}

func (s *DynamicStack[T]) Capacity() uint {
	return s.capacity
}

func (s *DynamicStack[T]) resize() {
	if s.IsFull() {
		s.capacity = s.capacity * 2
	} else {
		s.capacity = s.capacity / 2
	}
	newDate := make([]interface{}, s.capacity)
	copy(newDate, s.data)
	s.data = newDate
}

func (s *DynamicStack[T]) Push(data interface{}) error {
	if s.IsFull() {
		s.resize()
	}
	s.top++
	s.data[s.top] = data
	return nil
}

func (s *DynamicStack[T]) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	d := s.data[s.top]
	s.top--
	if s.Size() < s.capacity/2 {
		s.resize()
	}
	return d, nil
}

func (s *DynamicStack[T]) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.data[s.top], nil
}

func (s *DynamicStack[T]) Drain() {
	s.data = nil
	s.top = -1
}
