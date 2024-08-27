package main

import "fmt"

func isValid(s string) bool {

	stack := []rune{}
	parentises := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, item := range s {
		if item == '{' || item == '[' || item == '(' {
			stack = append(stack, item)
		}

		if item == '}' || item == ']' || item == ')' {
			itm := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if parentises[item] != itm {
				return false
			}
		}
	}
	return true

}

func main() {
	s := "()[]{}"

	fmt.Println("result : ", isValid(s))

}
