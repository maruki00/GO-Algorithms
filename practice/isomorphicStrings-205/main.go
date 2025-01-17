package main

import (
	"fmt"
	"sort"
)

func isIsomorphic(s string, t string) bool {
	sPattern, tPattern := map[uint8]int{}, map[uint8]int{}
	for index := range s {
		if sPattern[s[index]] != tPattern[t[index]] {
			return false
		} else {
			sPattern[s[index]] = index + 1
			tPattern[t[index]] = index + 1
		}
	}

	return true
}

func main() {
	fmt.Println("result : ", isIsomorphic("egg", "add"))
}
