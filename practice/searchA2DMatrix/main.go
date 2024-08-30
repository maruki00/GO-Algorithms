package main

import "fmt"

// func OneDSearch(nums []int, target int) bool {
// 	l, h := 0, len(nums)-1
// 	for l < h {
// 		middle := (l + h) / 2
// 		if nums[middle] == target {
// 			return true
// 		}
// 		if nums[l] < target {
// 			l = middle + 1
// 		} else {
// 			h = middle - 1
// 		}
// 	}
// 	return false

// }

// func searchMatrix(matrix [][]int, target int) bool {
// 	lr := 0
// 	hr := len(matrix) - 1

// 	for lr < hr {
// 		middlr := (lr + hr) / 2

// 		if OneDSearch(matrix[middlr], target) {
// 			return true
// 		}

// 		if OneDSearch(matrix[lr], target) {
// 			return true

// 		}
// 		if matrix[lr][0] < target {
// 			lr = middlr + 1
// 			continue
// 		}

// 		if OneDSearch(matrix[hr], target) {
// 			return true
// 		}
// 		if matrix[hr][0] > target {
// 			hr = middlr + 1
// 			continue
// 		}

// 	}

// 	return false
// }

func OneDSearch(nums []int, target int) bool {
	l, h := 0, len(nums)-1
	for l <= h {
		middle := (l + h) / 2
		if nums[middle] == target {
			return true
		}
		if nums[l] < target {
			l = middle + 1
		} else {
			h = middle - 1
		}
	}
	return false

}

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 1 {
		return OneDSearch(matrix[0], target)
	}
	lr := 0
	hr := len(matrix) - 1
	lstIndex := len(matrix[0]) - 1
	for lr <= hr {
		middlr := (lr + hr) / 2
		if matrix[middlr][0] <= target && matrix[middlr][lstIndex] >= target {
			return OneDSearch(matrix[middlr], target)
		}
		if matrix[middlr][0] < target {
			lr = middlr + 1
			continue
		} else {
			hr = middlr - 1
			continue
		}
	}
	return false
}

func main() {

	nums := [][]int{
		//{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
		//{1}, {3},
		{-8,-7,-5,-3,-3,-1,1},{2,2,2,3,3,5,7},{8,9,11,11,13,15,17},{18,18,18,20,20,20,21},{23,24,26,26,26,27,27},{28,29,29,30,32,32,34}
	}

	fmt.Println("result : ", searchMatrix(nums, 1))

}
