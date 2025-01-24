package main

import "fmt"

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[len(nums)-1]
	left := []int{}
	right := []int{}

	for _, n := range nums[:len(nums)-1] {
		if n > pivot {
			right = append(right, n)
			continue
		}
		left = append(left, n)
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

func main() {
	nums := []int{1, 2, 8, 4, 2, 8, 9, 1}
	fmt.Println(QuickSort(nums))
}
