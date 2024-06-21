package stack

import "errors"

type DynamicStack struct {
	top      int
	capacity uint
	data     []interface{}
}

func (s *DynamicStack) Init(capacity uint) *DynamicStack {
	s.top = -1
	s.capacity = capacity
	s.data = make([]interface{}, capacity)
	return s
}

func NewDynamicStack(capacity uint) *DynamicStack {
	return new(DynamicStack).Init(capacity)
}

func (s *DynamicStack) IsEmpty() bool {
	return s.top == -1
}

func (s *DynamicStack) IsFull() bool {
	return s.top == int(s.capacity)-1
}

func (s *DynamicStack) Size() uint {
	return uint(s.top + 1)
}

func (s *DynamicStack) Capacity() uint {
	return s.capacity
}

func (s *DynamicStack) resize() {
	if s.IsFull() {
		s.capacity = s.capacity * 2
	} else {
		s.capacity = s.capacity / 2
	}
	newDate := make([]interface{}, s.capacity)
	copy(newDate, s.data)
	s.data = newDate
}

func (s *DynamicStack) Push(data interface{}) error {
	if s.IsFull() {
		s.resize()
	}
	s.top++
	s.data[s.top] = data
	return nil
}

func (s *DynamicStack) Pop() (interface{}, error) {
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

func (s *DynamicStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.data[s.top], nil
}

func (s *DynamicStack) Drain() {
	s.data = nil
	s.top = -1
}
