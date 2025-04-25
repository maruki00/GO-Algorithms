package main

import "fmt"

func subsets(nums []int) [][]int {
	output := [][]int{{}}
	for _, num := range nums {
		newSubsets := [][]int{}
		for _, curr := range output {
			newSubset := append([]int{num}, curr...)
			newSubsets = append(newSubsets, newSubset)
		}
		output = append(output, newSubsets...)
	}
	return output
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("result : ", subsets(nums))
}
