package main

import "fmt"

var stack []int

func largestRectangleArea(heights []int) int {
	stack = stack[:0]
	maxArea := 0
	n := len(heights)

	for i := 0; i <= n; i++ {
		h := 0
		if i < n {
			h = heights[i]
		}
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			hIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := heights[hIndex]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}

func main() {
	hs := []int{2, 1, 5, 6, 2, 3}
	fmt.Println("result : ", largestRectangleArea(hs))
}
