package main

import "fmt"

/*
	9. Implement quick sort.
*/

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[len(nums)-1]
	left, right := []int{}, []int{}
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
	nums := []int{1, 2, 3, 56, 3, 6, 23, 34, 235, 234, 346, 345234, 5, 6756, 34, 23452, 74567, 2345, 3245, 563456, 3456, 467, 3}
	fmt.Println("result : ", QuickSort(nums))
}
