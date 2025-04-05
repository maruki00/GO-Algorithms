package main

import "fmt"

func CanPlaceFlowers(nums []int, n int) bool {

	for i := 0; i < len(nums); i++ {
		if n <= 0 {
			break
		}
		if len(nums) == 1 && nums[0] == 0 {
			return true
		}
		if nums[i] == 1 {
			continue
		}
		if i == 0 && nums[i+1] != 1 {
			nums[i] = 1
			n--
			continue
		}

		if i == len(nums)-1 && nums[i-1] != 1 {
			nums[i] = 1
			n--
			continue
		}
		if i > 0 && nums[i-1] != 1 && nums[i+1] != 1 {
			nums[i] = 1
			n--
		}
	}
	return n <= 0
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 0, 0, 1}
	n := 2
	fmt.Println("result : ", CanPlaceFlowers(flowerbed, n))
}
