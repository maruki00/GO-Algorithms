package main

import (
	"fmt"
	"math"
)

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	if n < 10 {
		return false
	}
	num := 0
	for n > 0 {
		num = num*10 + int(math.Pow(float64(num%10), float64(2)))
		n /= 10
	}
	fmt.Println(n, num)
	return isHappy(num)
}

func main() {
	fmt.Println("result : ", isHappy(19))
}
