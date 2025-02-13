package main

import "fmt"

func tribonacci(n int) int {
	tracking := make([]int, n+10)
	tracking[0] = 0
	tracking[1] = 1
	tracking[2] = 1
	for i := 0; i <= n; i++ {
		tracking[i+3] = tracking[i] + tracking[i+1] + tracking[i+2]
	}
	return tracking[n]
}

func main() {
	n := 25
	fmt.Println("result : ", tribonacci(n))
}
