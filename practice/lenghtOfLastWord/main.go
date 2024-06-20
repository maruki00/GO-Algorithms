package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	res := strings.Split(s, " ")

	return len(res[len(res)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
}
