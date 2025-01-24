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

func QuickSort_(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}
	left := []int{}
	right := []int{}

	pivot := arr[len(arr)-1]
	for _, item := range arr[:len(arr)-1] {
		if item > pivot {
			right = append(right, item)
			continue
		}
		left = append(left, item)
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

func main() {
	nums := []int{1, 2, 8, 4, 2, 8, 9, 1}
	fmt.Println(QuickSort(nums))
}
