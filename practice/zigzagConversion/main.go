package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	var ret []string = make([]string, numRows)
	index := 0
	maxLen := len(s)
	run := true
	for run {
		i := 0
		for i = 0; i < numRows; i++ {
			if index >= maxLen {
				run = false
				break
			}
			fmt.Println("I : ", i)
			ret[i] = ret[i] + string(s[index])
			index++
		}

		for j := i - 2; j > 0; j-- {
			if index >= maxLen {
				run = false
				break
			}
			fmt.Println("J : ", j)
			ret[j] = ret[j] + string(s[index])
			index++
		}

	}

	return strings.Join(ret, "")

}

func main() {
	s := "PAYPALISHIRING"
	num := 3
	fmt.Println(convert(s, num))
}
