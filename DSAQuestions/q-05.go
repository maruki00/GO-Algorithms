package main

import "fmt"

/*
	5. Find the factorial of a number using recursion
*/

func fact(num int) int {
	if num <= 1 {
		return 1
	}
	return num * fact(num-1)
}

func main() {
	num := 10
	fmt.Println("result : ", fact(num))
}
