package main

import (
	"fmt"
)

func main() {
	items := []int{
		1, 6, 8, 4,
		3, 8, 0, 44,
		43, 23, 64, 23,
		676, 234, 5, 234,
	}

	searchFor := 32
	foundAt := -1

	for i, val := range items {
		if val == searchFor {
			foundAt = i
		}
	}

	if foundAt != -1 {
		fmt.Printf("The %d founded at index %d\n", searchFor, foundAt)
	} else {
		fmt.Printf("The %d is not found\n", searchFor)
	}
}
