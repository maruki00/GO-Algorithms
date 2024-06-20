package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	num1, _ := strconv.ParseInt(a, 2, 32)
	num2, _ := strconv.ParseInt(b, 2, 32)

	println(num1, num2)
	return "111"
}

func main() {

	fmt.Println(addBinary("11", "11"))
}
