package main

import (
	"fmt"
)

func main() {

	items := []int{1, 2, 3, 4, 5, 6}
	left := 0
	right := len(items) - 1
	ans := 0

	minItem := 0
	for left < right {
		minItem = min(items[left], items[right])
		ans = max(ans, ((right - left) * minItem))
		if items[left] > items[right] {
			right--
		} else {
			left++
		}
	}

	fmt.Print(items, left, right)
}
