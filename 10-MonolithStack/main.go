package main

import (
	"fmt"
)

// Monolithic stack to be reused across functions
var stack []int

func nextGreaterElement(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1 // Default to -1 if no greater element exists
	}

	// Clear the stack before using it
	stack = stack[:0]

	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[index] = nums[i]
		}
		stack = append(stack, i)
	}
	return result
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n) // Result array initialized with 0s

	// Clear the stack before using it
	stack = stack[:0]

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return result
}

func main() {
	nums := []int{2, 1, 5, 6, 2, 3}
	fmt.Println("Next Greater Element Input:", nums)
	fmt.Println("Next Greater Elements:", nextGreaterElement(nums))

	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println("Daily Temperatures Input:", temperatures)
	fmt.Println("Days to Warmer Temperature:", dailyTemperatures(temperatures))
}
