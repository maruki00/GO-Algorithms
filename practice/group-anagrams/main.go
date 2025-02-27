package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	// hashMap := make(map[string][]string)
	hashMap := make(map[[26]int][]string)
	result := make([][]string, 0)
	for _, j := range strs {
		// tmp := strings.Split(j, "")
		// sort.Strings(tmp)
		// key := strings.Join(tmp, "")
		// hashMap[key] = append(hashMap[key], j)

		tmp := [26]int{}
		for _, j := range j {
			tmp[int(j)-97] += 1
		}
		hashMap[tmp] = append(hashMap[tmp], j)

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
