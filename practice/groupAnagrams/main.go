package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	hashMap := make(map[int][]string)
	result := make([][]string, 0)
	for _, j := range strs {
		hashMap[sort.Strings(j)] = append(hashMap[tmp], j)
	}
	for _, i := range hashMap {
		result = append(result, i)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("result :", groupAnagrams(strs))
}
