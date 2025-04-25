package main

import "fmt"

func permute(nums []int) [][]int {

	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		result = append(result, append([]int{nums[i]}, append(nums[i+1:], nums[:i]...)...))
	}
	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
