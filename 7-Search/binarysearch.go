package main

import (
	"fmt"
)

func main() {
	items := []int{
		1, 5, 8, 10,
		23, 45, 56, 67,
		78, 89, 98, 100,
	}
	isFound := false
	target := 1001
	low := 0
	hight := len(items) - 1
	middle := (low + hight) / 2
	for hight >= low {
		if items[middle] == target {
			isFound = true
			break
		} else if items[middle] > target {
			hight = middle - 1
		} else {
			low = middle + 1
		}
		middle = ((hight + low) / 2)
	}
	if isFound == true {
		fmt.Printf("Found at %d\n", middle)
	} else {
		fmt.Printf("Not Found\n")
	}
}
