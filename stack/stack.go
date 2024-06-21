package stack

type Stack interface {
	Push(data interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	Drain()
	IsEmpty() bool
	IsFull() bool
	Size() uint
	Capacity() uint
}
