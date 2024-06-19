package stack

import (
	"errors"
)

type SimpleStack struct {
	top      int
	capacity uint
	data     []interface{}
}

func (s *SimpleStack) Init(capacity uint) *SimpleStack {
	s.top = -1
	s.capacity = capacity
	s.data = make([]interface{}, capacity)
	return s
}

func NewSimpleStack(capacity uint) *SimpleStack {
	return new(SimpleStack).Init(capacity)
}

func (s *SimpleStack) IsEmpty() bool {
	return s.top == -1
}

func (s *SimpleStack) IsFull() bool {
	return s.top == int(s.capacity)-1
}

func (s *SimpleStack) Size() uint {
	return uint(s.top + 1)
}

func (s *SimpleStack) Push(data interface{}) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}
	s.top++
	s.data[s.top] = data
	return nil
}

func (s *SimpleStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	d := s.data[s.top]
	s.top--
	return d, nil
}

func (s *SimpleStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.data[s.top], nil
}

func (s *SimpleStack) Drain() {
	s.data = nil
	s.top = -1
}
