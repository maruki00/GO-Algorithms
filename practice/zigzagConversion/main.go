package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	var ret []string = make([]string, numRows)
	row := 0
	inc := true
	for _, item := range s {

		ret[row] = ret[row] + string(item)

		if row == 0 {
			inc = true
			row++
		} else if row == numRows-1 {
			inc = false
			row--
		}

	}
	s = ""
	for _, res := range ret {
		s += res
	}
	return s
}

func main() {
	s := "PAYPALISHIRING"
	num := 3
	fmt.Println(convert(s, num))
}
