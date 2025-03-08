package main

import (
	"fmt"
)

func carFleet(target int, position []int, speed []int) int {
	res := 0
	current := 0.0
	stack := make([]float64, target)

	for i, pos := range position {
		stack[pos] = float64(target-pos) / float64(speed[i])
	}

	for i := target - 1; i >= 0; i-- {
		if stack[i] > current {
			current = stack[i]
			res++
		}
	}

	return res
}

func main() {
	target := 12
	position := []int{12, 8, 0, 5, 3}
	speed := []int{0, 4, 1, 1, 3}
	fmt.Println("result : ", carFleet(target, position, speed))
}
