package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	//s = strings.Split(s, ' ')

	fmt.Print(s)
}

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
}
