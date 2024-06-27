package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	items []int
}

func (obj *Stack) pop() int {
	l := len(obj.items)
	fmt.Println("Pop Index: ", l, obj.items)
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
			x = s.pop()
			y = s.pop()
			s.push(x + y)
			fmt.Println("+", y, x)
		case "-":
			x = s.pop()
			y = s.pop()
			s.push(x - y)
			fmt.Println("-", y, x)
		case "*":
			x = s.pop()
			y = s.pop()
			s.push(x * y)
			fmt.Println("*", y, x)
		case "/":
			x = s.pop()
			y = s.pop()
			s.push(x / y)
			fmt.Println("/", y, x)
		default:
			if item, err := strconv.Atoi(token); err == nil {
				s.push(item)
			}
			fmt.Println("default", y, x)
		}
		fmt.Println("token: ", token, s.items)
	}

	return s.pop()
}

func main() {
	token := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(token))

}
