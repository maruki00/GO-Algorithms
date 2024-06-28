package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	var stack []int = make([]int, len(temperatures))
	//var ans []int = make([]int, len(temperatures))
	l := 0
	for i := 0; i < len(temperatures); i++ {

		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			l = len(stack)
			last := stack[l-1]
			stack = stack[:l-1]
			temperatures[last] = i - last
		}
		stack = append(stack, i)
	}
	return temperatures
}

func main() {

	temps := []int{99, 99, 99, 99, 99, 99, 99, 99, 100, 9, 9, 9, 100}

	fmt.Println("result : ", dailyTemperatures(temps))
}
