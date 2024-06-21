package problem

import (
	"ds-by-go/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

// find kth from end
func TestFindFromEnd(t *testing.T) {
	l := list.SingleLinkedList{}
	l.InsertBeginning(1)
	l.InsertBeginning(2)
	l.InsertBeginning(3)
	l.InsertBeginning(4)
	l.InsertBeginning(5)

	assert.Equal(t, 2, l.FindKthFromEnd(2))
}
