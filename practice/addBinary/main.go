package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	num1, _ := strconv.ParseInt(a, 2, 1024)
	num2, _ := strconv.ParseInt(b, 2, 1024)

	return strconv.FormatInt(num1+num2, 2)
}

func main() {

	fmt.Println(addBinary("11", "1"))
}
