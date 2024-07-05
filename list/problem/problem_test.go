package problem

import (
	"ds-by-go/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

// find kth from end
func TestFindFromEnd(t *testing.T) {
	ll := list.NewSingleLinkedList([]interface{}{1, 2, 3, 4, 5})
	assert.Equal(t, 4, ll.FindKthFromEnd(2))
}
