package main

import (
	"fmt"
	"strings"
)

func removeStarts(s string) string {
	stack := []string{}
	for _, i := range s {
		if string(i) == "*" {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, string(i))
		}
	}

	return strings.Join(stack, "")

}

func main() {
	fmt.Println("result : ", removeStarts("stack**leet*code"))
}
