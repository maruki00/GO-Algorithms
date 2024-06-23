package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var ret []string = make([]string, numRows)
	row := 0
	inc := true
	for _, item := range s {
		if row == 0 {
			inc = true
		} else if row == numRows-1 {
			inc = false
		}
		ret[row] = ret[row] + string(item)
		if inc {
			row++
		} else {
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
	s := "AB" //"PAYPALISHIRING"
	num := 1
	fmt.Println(convert(s, num))
}
