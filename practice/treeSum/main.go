package main

import "fmt"

func threeSum(nums []int) [][]int {
	l := len(nums) - 1
	hashMap := make(map[int][]int)
	var result [][]int = [][]int{}
	for j, i := range nums {
		hashMap[i] =  append(hashMap[i], j)
	}
	for index, _ := range nums {
		left, right := index+1, l
		for left < l {

			/*if nums[index]+nums[left]+nums[right] == 0 {
				result = append(result, []int{nums[index], nums[left], nums[right]})
				break
			}*/

			lNum := nums[index] + nums[left]
			if _, ok := hashMap[0-lNum]; ok {
				found := false
				for _,m := range hashMap[0-lNum]{
					if m > left{
						result = append(result, []int{nums[index], nums[left], 0 - lNum})
						found = true
						break
				}
				if found {
					break
				}
			}

			/*rNum := nums[index] + nums[right]
			if _, ok := hashMap[0-rNum]; ok {
				result = append(result, []int{nums[index], nums[right], 0 - rNum})
				break
			}
			right--
			left++
			*/
			left ++
		}

	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("Result : ", threeSum(nums))
}
