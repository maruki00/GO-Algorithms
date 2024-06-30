package main

import "fmt"

func maxArea(height []int) int {
	l := len(height) - 1
	left, right := height[0], height[l]
	maxA := 0
	for left < right {
		maxA = min(maxA, (right-left)*min(right, left))
		if left < right {
			left++
		} else if left > right {
			right--
		} else {
			left++
			right--
		}
	}
	return maxA
}

func main() {
	hieghts := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Result : ", maxArea(hieghts))
}
