package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	res := strings.Split(s, " ")

	fmt.Print(res[len(res)-1])
	return 0
}

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
}
