package main

import "fmt"

func climbStairs(n int) int {
	left := 1
	right := 1
	for i := 2; i <= n; i++ {
		right, left = (left + right), right
	}
	return right
}

func main() {
	n := 5
	fmt.Println("Stairs : ", climbStairs(n))
}
