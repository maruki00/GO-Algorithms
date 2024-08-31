package main

import "fmt"

func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fib2(n-1) + fib2(n-2)
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	fibMem := [2]int{0, 1}
	for i := 2; i < n; i++ {
		tmp := fibMem[0] + fibMem[1]
		fibMem[0] = fibMem[1]
		fibMem[1] = tmp
	}
	return fibMem[0] + fibMem[1]
}

func main() {
	fmt.Println("result : ", fib(10))
}
