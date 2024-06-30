package main

import (
	"fmt"
	"sort"
)

func isExists(nums [][]int, num [3]int) bool {
	for _, x := range nums {
		if num[0] == x[0] && num[1] == x[1] && num[2] == x[2] {
			return true
		}
	}
	return false
}
func threeSum(nums []int) [][]int {
	l := len(nums) - 1
	hashMap := make(map[int][]int)
	var result [][]int = [][]int{}
	for j, i := range nums {
		hashMap[i] = append(hashMap[i], j)
	}
	fmt.Println(hashMap)
	for index, _ := range nums {
		left := index + 1
		for left < l {

			lNum := nums[index] + nums[left]
			if _, ok := hashMap[0-lNum]; ok {
				for _, m := range hashMap[0-lNum] {
					if m > left {
						tmp := [3]int{nums[index], nums[left], 0 - lNum}
						sort.Ints(tmp)
						if isExists(result, tmp) == false {
							result = append(result, tmp[:])
							break
						}

					}
				}
			}

			left++
		}

	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("Result : ", threeSum(nums))
}
