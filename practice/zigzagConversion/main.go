package main

import (
	"fmt"
	"strings"
)

// func convert(s string, numRows int) string {
// 	if numRows == 1 {
// 		return s
// 	}
// 	var ret []string = make([]string, numRows)
// 	row := 0
// 	inc := true
// 	for _, item := range s {
// 		if row == 0 {
// 			inc = true
// 		} else if row == numRows-1 {
// 			inc = false
// 		}
// 		ret[row] = ret[row] + string(item)
// 		if inc {
// 			row++
// 		} else {
// 			row--
// 		}
// 	}
// 	return strings.Join(ret, "")
// 	s = ""
// 	for _, res := range ret {
// 		s += res
// 	}
// 	return s
// }

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var ret []string = make([]string, numRows)
	row := 0
	inc := true
	for _, item := range s {
		ret[row] = ret[row] + string(item)

		if row == 0 || inc {
			row++
			inc = true
		} else if row == numRows-1 || !inc {
			row--
			inc = false
		}
	}
	return strings.Join(ret, "")

}

func main() {
	s := "AB" //"PAYPALISHIRING"
	num := 1
	fmt.Println(convert(s, num))
}
