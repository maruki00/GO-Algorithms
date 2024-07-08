package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {

	hashMap := make(map[string][]string)
	result := make([][]string, 0)
	for _, j := range strs {
		tmp := strings.Split(j, "")
		sort.Strings(tmp)
		key := strings.Join(tmp, "")
		hashMap[key] = append(hashMap[key], j)
	}
	Join()
	for _, i := range hashMap {
		result = append(result, i)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("result :", groupAnagrams(strs))
}
