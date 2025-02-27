package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		val := (a & b) << 1
		a = a ^ b
		b = val
	}
	return a
}

func main() {
	fmt.Println("result : ", getSum(1, 5))
}
