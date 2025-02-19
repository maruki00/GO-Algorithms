package main

import (
	"fmt"
	"math"
	"sort"
)

func thirdMax_optimized(nums []int) int {
	maxes := [3]int{
		math.MinInt,
		math.MinInt,
		math.MinInt,
	}

	for _, num := range nums {
		if maxes[0] == num || maxes[1] == num || maxes[2] == num {
			continue
		}
		if num > maxes[0] || num == maxes[0] {
			maxes[0], maxes[1], maxes[2] = num, maxes[0], maxes[1]
			continue
		}
		if num > maxes[1] || num == maxes[2] {
			maxes[1], maxes[2] = num, maxes[1]
			continue
		}
		if num > maxes[2] || num == maxes[2] {
			maxes[2] = num
			continue
		}
	}
	if maxes[2] == math.MinInt {
		return maxes[0]
	}
	return maxes[2]
}

func SetFromInts(nums []int) []int {
	tracking := make(map[int]bool)
	s := []int{}
	for _, i := range nums {
		if _, ok := tracking[i]; ok {
			continue
		}
		tracking[i] = true
		s = append(s, i)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	return s
}

func Len(s []int) int {
	return len(s)
}

func At(s []int, i int) int {
	if Len(s) == 0 {
		return 0
	}
	return s[i]
}

func thirdMax(nums []int) int {
	fmt.Println(nums)
	s := SetFromInts(nums)

	if Len(s) > 2 {
		return At(s, 2)
	}

	return At(s, 0)
}

func main() {
	nums := []int{2, 2, 3, 1}
	fmt.Println("result : ", thirdMax(nums))
}
