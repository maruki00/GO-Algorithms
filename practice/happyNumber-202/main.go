package main

import (
	"fmt"
	"math"
)

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	if n < 10 && n%2 == 0 {
		return false
	}
	num := 0
	for n > 0 {
		num += int(math.Pow(float64(n%10), float64(2)))
		n /= 10
	}
	return isHappy(num)
}

func main() {
	fmt.Println("result : ", isHappy(19))
}
