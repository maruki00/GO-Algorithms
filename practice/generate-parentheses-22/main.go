package main

import (
	"fmt"
	"strings"
)

type Stack []string

func (s *Stack) Pop() string {
	old := *s
	lst := old[len(old)-1]
	old = old[:len(old)-1]
	*s = old
	return lst
}
func (s *Stack) Push(b string) {
	*s = append(*s, b)
}
func generateParenthesis(n int) []string {
	stack := &Stack{}
	res := []string{}
	var tracking func(nopen, nclose int)
	tracking = func(nopen, nclose int) {
		fmt.Println(*stack)
		if nopen == n && n == nclose {
			res = append(res, strings.Join(*stack, ""))
			return
		}

		if nopen < n {
			stack.Push("(")
			tracking(nopen+1, nclose)
			_ = stack.Pop()
		}
		if nopen > nclose {
			stack.Push(")")
			tracking(nopen, nclose+1)
			_ = stack.Pop()
		}
	}
	tracking(0, 0)
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
