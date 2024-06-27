package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	items []int
}

func (obj *Stack) pop() int {
	l := len(obj.items) - 1
	item := obj.items[l]
	obj.items = obj.items[:l]
	return item
}

func (obj *Stack) push(val int) {
	obj.items = append(obj.items, val)
}

func evalRPN(tokens []string) int {
	s := &Stack{}

	x := 0
	y := 0

	for _, token := range tokens {

		switch token {
		case "+":
			y = s.pop()
			x = s.pop()
			s.push(x + y)
		case "-":
			y = s.pop()
			x = s.pop()
			s.push(x - y)
		case "*":
			x = s.pop()
			y = s.pop()
			s.push(x * y)
		case "/":
			y = s.pop()
			x = s.pop()
			s.push(x / y)
		default:
			if item, err := strconv.Atoi(token); err == nil {
				s.push(item)
			}
		}
	}

	return s.pop()
}

func main() {
	token := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(token))

}
