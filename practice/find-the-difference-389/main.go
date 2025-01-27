package main

import "strings"

func findTheDifference(s string, t string) byte {
	sMap := make(map[byte]int, 0)
	tMap := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	for key, val := range sMap {
		num, ok := tMap[key]
		if ok && val == num {
			delete(sMap, key)
			delete(tMap, key)
			continue
		}
	}
	result := []byte{}
	for key, _ := range sMap {
		result = append(result, key)
	}
	for key, _ := range tMap {
		result = append(result, key)
	}
	return result[0]
}

func main() {

}
