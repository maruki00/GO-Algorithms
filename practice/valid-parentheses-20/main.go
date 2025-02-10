package main

import "fmt"

type Stack struct {
	items []byte
}

func (s *Stack) Push(item byte) {
	s.items = append(s.items, item)
}
func (s *Stack) Pop() byte {
	l := len(s.items) - 1
	lst := s.items[l]
	s.items = s.items[:l]
	return lst
}
func (s *Stack) Empty() bool {
	return len(s.items) == 0
}
func isValid(s string) bool {
	items := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := &Stack{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack.Push(s[i])
			continue
		}
		if stack.Empty() {
			return false
		}
		itm := stack.Pop()
		if s[i] != items[itm] {
			return false
		}
	}
	return stack.Empty()
}

func main() {
	s := "]"
	fmt.Println(isValid(s))
}
