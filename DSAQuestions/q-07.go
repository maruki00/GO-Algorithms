package main

import (
	"fmt"
	"math"
)

/*
	7. Find the largest/smallest element in an array.
*/

func main() {
	nums := []int{1, 2, 3, 65, 3, 6, 7, 32, 68, 3, 7, 3, 78, 98, 3}
	minItem, maxItem := math.MaxInt, math.MinInt
	for _, j := range nums {
		minItem = min(minItem, j)
		maxItem = max(maxItem, j)
	}
	fmt.Println("min: ", minItem, "\nmax: ", maxItem)
}
