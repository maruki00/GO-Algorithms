package main

import "math"

func kidsWithCandies(candies []int, extraCandies int) []bool {

	maxCandies := math.MinInt

	for i := range candies {
		maxCandies = max(maxCandies, candies[i])
	}
	result := make([]bool, len(candies))
	// In the given code snippet, `candie` is a variable that represents the number of candies for
	// a particular child in the `candies` slice. It is used in the loop to iterate over each
	// child's candies and check if adding `extraCandies` to the current child's candies would make
	// them have the maximum number of candies among all children.
	for i, candie := range candies {
		result[i] = candie+extraCandies >= maxCandies
	}
	return result
}
