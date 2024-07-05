package problem

import (
	"ds-by-go/stack"
)

type pair struct {
	open  rune
	close rune
}

var pairs = []pair{
	{'(', ')'},
	{'[', ']'},
	{'{', '}'},
}

// CheckStringLiteral PROBLEM_01 checking balancing of symbols/parentheses
func CheckStringLiteral(str string) bool {
	stack := stack.NewDynamicStack[string](1)
	for _, r := range str {
		for _, p := range pairs {
			temp, _ := stack.Peek()
			if r == p.open {
				stack.Push(r)
				break
			} else if r == p.close && stack.IsEmpty() {
				return false
			} else if r == p.close && temp == p.open {
				stack.Pop()
				break
			} else if r == p.close && temp != p.open {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
