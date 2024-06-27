package main

import "fmt"

type Stack struct {
	items []float32
}

func (obj *Stack) pop() float32 {
	l := len(obj.items) - 1
	item := obj.items[l]
	obj.items = obj.items[:l]
	return item
}

func (obj *Stack) ppush(val float32) {
	obj.items = append(obj.items, val)
}

func evalRPN(tokens []string) int {
	s := &Stack{}
	fmt.Println(s)

	return 0
}

func main() {
	token := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(token))
}
