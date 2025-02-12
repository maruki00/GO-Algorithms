package main

import (
	"container/list"
	"fmt"
	"unicode"
)

func clearDigits(s string) string {
	stack := list.New()
	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) {
			e := stack.Front()
			stack.Remove(e)
			continue
		}
		stack.PushFront(string(s[i]))
	}
	result := ""
	for stack.Len() > 0 {
		e := stack.Back()
		result += e.Value.(string)
		stack.Remove(e)
	}
	return result
}

func main() {
	s := "cb34aa"
	fmt.Println("result : ", clearDigits(s))
}
