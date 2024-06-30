package main

import "fmt"

func maxArea(height []int) int {
	l := len(height) - 1
	left, right := 0, l
	maxA := 0
	for left < right {
		maxA = max(maxA, (right-left)*min(height[right], height[left]))
		if height[left] < height[right] {
			left++
		} else if height[left] > height[right] {
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
