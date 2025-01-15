package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	lS, lt := len(s), len(t)
	if lS != lt {
		return false
	}

	sChrs := make(map[byte]int, 0)
	tChrs := make(map[byte]int, 0)
	for i := 0; i < lS; i++ {
		sChrs[s[i]]++
		tChrs[t[i]]++
	}
	if len(sChrs) != len(tChrs) {
		return false
	}

	for k, v := range sChrs {
		if val, ok := tChrs[k]; !ok || val != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("result : ", isIsomorphic("egg", "add"))
}
