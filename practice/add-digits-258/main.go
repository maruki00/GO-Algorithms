package main

import "fmt"

func addDigits(num int) int {
	var solution func(n int) int
	solution = func(n int) int {
		if n <= 9 {
			return n
		}
		result := 0
		for n > 0 {
			result += +int(n % 10)
			n /= 10
		}
		return solution(result)
	}
	return solution(num)
}

func main() {
	fmt.Println("result : ", addDigits(38))
}
