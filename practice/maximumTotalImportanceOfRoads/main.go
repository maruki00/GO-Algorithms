package main

import (
	"fmt"
	"sort"
)

func maximunImportance(n int, roads [][]int) int64 {
	hashSet := make(map[int]int)
	for _, i := range roads {
		hashSet[i[0]]++
		hashSet[i[1]]++
	}
	keys := []int{}
	for _, i := range hashSet {
		keys = append(keys, i)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return hashSet[keys[i]] < hashSet[keys[j]]
	})

	fmt.Println("Result : ", hashSet)
	return 0
}

func main() {
	roads := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{0, 2},
		{1, 3},
		{2, 4},
	}
	fmt.Println(maximunImportance(5, roads))
}
