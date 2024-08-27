package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	Items []int
}

func (obj *Stack) push(item int) {
	obj.Items = append(obj.Items, item)
}

func (obj *Stack) Empty() bool {
	return len(obj.Items) == 0
}

func (obj *Stack) pop() int {
	fmt.Println(" poping ... : ", obj.Items)
	l := len(obj.Items)
	if l == 0 {
		return 0
	}
	item := obj.Items[l-1]
	obj.Items = obj.Items[:l-1]
	return item
}

func calPoints(operations []string) int {
	stack := &Stack{}

	for _, item := range operations {
		if item == "C" && !stack.Empty() {
			_ = stack.pop()
			continue
		}

		if item == "D" && !stack.Empty() {
			num := stack.pop()
			stack.push(num)
			stack.push(num * 2)
			continue
		}

		if item == "+" {
			first, second := 0, 0
			if !stack.Empty() {
				first = stack.pop()
			}
			if !stack.Empty() {
				second = stack.pop()
			}
			stack.push(second)
			stack.push(first)
			stack.push(first + second)
			continue
		}

		num, err := strconv.Atoi(item)
		if err != nil {
			continue
		}
		stack.push(int(num))

	}

	res := 0
	for _, val := range stack.Items {
		res += val
	}

	return res

}

func main() {
	strr := []string{"5", "2", "C", "D", "+"}
	fmt.Println("result : ", calPoints(strr))
}
