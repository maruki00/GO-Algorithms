package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Map := make(map[byte]int)
	s2Map := make(map[byte]int)
	for i := range s1 {
		s1Map[s1[i]]++
	}
	start, end := 0, len(s1)-1
	for k := start; k <= end; k++ {
		s2Map[s2[k]]++
	}
	found := false
	for end < len(s2) {
		found = true
		for k, v := range s1Map {
			if s2Map[k] != v {
				found = false
				break
			}
		}
		if found {
			return true
		}
		start++
		end++
		if end >= len(s2) {
			break
		}
		s2Map[s2[start-1]]--
		s2Map[s2[end]]++
	}
	return false
}

func main() {
	// s1 := "ab"
	// s2 := "eidbaooo"

	s1 := "adc"
	s2 := "dcda"

	fmt.Println("result : ", checkInclusion(s1, s2))
}
