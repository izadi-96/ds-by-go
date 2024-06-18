package list

type LinkedList interface {
	InsertBeginning(data interface{})
	InsertEnd(data interface{})
	Insert(position int, data interface{}) error
	DeleteFirst() (interface{}, error)
	DeleteLast() (interface{}, error)
	DeleteAt(position int) (interface{}, error)
	GetAtPosition(position int) (interface{}, error)
	Display() error
	Length() int
}
