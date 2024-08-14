package main

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}

	for i := 0; i < len(nums); i++ {
		j := i
		for j < len(nums)-1 {
			if nums[i]+1 == nums[j+1] {
				j++
			} else {
				break
			}
		}

		if i == j {
			res = append(res, fmt.Sprintf("%d", nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j]))
		}
		i = j
	}

	return res
}

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println("result : ", summaryRanges(nums))
}
