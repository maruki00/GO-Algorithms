package main

import (
	"fmt"
	"sort"
)

// func maximunImportance(n int, roads [][]int) int64 {
// 	hashSet := make(map[int][]int, n)
// 	result := int64(0)

// 	for _, i := range roads {
// 		if len(hashSet[i[0]]) == 0 {
// 			hashSet[i[0]] = []int{0, 0}
// 		}
// 		if len(hashSet[i[1]]) == 0 {
// 			hashSet[i[1]] = []int{0, 0}
// 		}

// 		hashSet[i[0]][0] = hashSet[i[0]][0] + 1
// 		hashSet[i[1]][0] = hashSet[i[1]][0] + 1
// 	}

// 	keys := make([]int, 0, n)
// 	for i := range hashSet {
// 		keys = append(keys, i)
// 	}
// 	sort.SliceStable(keys, func(i, j int) bool {
// 		if hashSet[keys[i]][0] == hashSet[keys[j]][0] {
// 			return i < j
// 		} else {
// 			return hashSet[keys[i]][0] < hashSet[keys[j]][0]
// 		}
// 	})

// 	for i,j := range keys {

// 	}

// 	fmt.Println(keys)

// 	// res := make(map[int][]int, 5)
// 	// for i := 0; i < len(keys); i++ {
// 	// 	result += int64()
// 	// 	res[keys[i]] = append(res[keys[i]], []int{hashSet[keys[i]], i + 1}...)
// 	// }
// 	// fmt.Println(hashSet, keys, res)
// 	return result
// }

func maximumImportance(n int, roads [][]int) int64 {
	items := make([]int, n)
	for _, e := range roads {
		items[e[0]]++
		items[e[1]]++
	}

	fmt.Println(items)
	sort.Ints(items)
	fmt.Println(items)

	var res int64 = 0
	for i := 0; i < n; i++ {
		res += int64(i+1) * int64(items[i])
	}
	return res
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
	fmt.Println(maximumImportance(5, roads))
}
