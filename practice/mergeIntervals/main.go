package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}

	tmp := []int{}
	for _, interval := range intervals {
		if len(tmp) == 0 {
			tmp = interval
			continue
		}
		if tmp[1] >= interval[0] {
			tmp[1] = max(interval[0], interval[1], tmp[1])
			tmp[0] = min(interval[0], interval[1], tmp[0])
			continue
		}

		res = append(res, tmp)
		tmp = interval

	}
	if len(tmp) != 0 {
		res = append(res, tmp)
	}

	return res
}

func main() {

	intervals := [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18}, {1, 4},
	}

	fmt.Println("result : ", merge(intervals))

}
