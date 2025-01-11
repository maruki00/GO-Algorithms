package main

import (
	"fmt"
)

func equals(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	hashMap := make(map[byte]int)
	for i, _ := range s1 {
		hashMap[s1[i]]++
		hashMap[s2[i]]--
	}
	for _, item := range hashMap {
		if item != 0 {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	l := len(s1)
	for i := 0; i < len(s2)-l+1; i++ {
		if equals(s2[i:i+l], s1) {
			return true
		}
	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidboaoo" //"eidbaooo"
	fmt.Println("result : ", checkInclusion(s1, s2))
}
