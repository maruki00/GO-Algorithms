package main

import "fmt"

func OneDSearch(nums []int, target int) bool {
	l, h := 0, len(nums)-1
	for l <= h {
		middle := (l + h) / 2
		if nums[middle] == target {
			return true
		}
		if nums[middle] < target {
			l = middle + 1
		} else {
			h = middle - 1
		}
	}
	return false

}

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}
	if len(matrix) == 1 {
		return OneDSearch(matrix[0], target)
	}

	firstIndex, lastIndex := 0, len(matrix[0])-1

	lr, hr := 0, len(matrix)-1

	for lr <= hr {
		middle := (lr + hr) / 2
		if matrix[middle][firstIndex] <= target && matrix[middle][lastIndex] >= target {
			return OneDSearch(matrix[middle], target)
		}
		if matrix[lr][firstIndex] <= target && matrix[lr][lastIndex] >= target {
			return OneDSearch(matrix[lr], target)
		}

		if matrix[hr][firstIndex] <= target && matrix[hr][lastIndex] >= target {
			return OneDSearch(matrix[hr], target)
		}

		if matrix[middle][lastIndex] > target {
			hr = middle - 1
		} else {
			lr = middle + 1
		}
	}
	return false
}
func main() {

	nums := [][]int{
		// {1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
		//{1}, {3},
		//{-8,-7,-5,-3,-3,-1,1},{2,2,2,3,3,5,7},{8,9,11,11,13,15,17},{18,18,18,20,20,20,21},{23,24,26,26,26,27,27},{28,29,29,30,32,32,34}
		{1}, {3},
	}

	fmt.Println("result : ", searchMatrix(nums, 1))

}
