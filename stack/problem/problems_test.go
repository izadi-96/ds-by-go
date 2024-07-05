package problem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckStringLiterals(t *testing.T) {
	testCases := []struct {
		input  string
		result bool
	}{
		{input: "()", result: true},
		{input: "(]", result: false},
		{input: "", result: true},
		{input: "(())()[]", result: true},
		{input: ")", result: false},
	}

	for _, tc := range testCases {
		result := CheckStringLiteral(tc.input)
		assert.Equal(t, tc.result, result)
	}

}
