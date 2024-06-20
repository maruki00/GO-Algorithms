package main

import "fmt"

func getMinLen(s []string) int {
	minLen := 0
	for i := range s {
		if len(i) < minLen {
			minLen = len(i)
		}
	}
	return minLen
}

func longestCommonPrefix(strs []string) string {
	var prifix string = ""
	minLenght := getMinLen(strs)
	for i := 0; i < minLenght; i++ {
		tmp := strs[0][i]
		isOk := true
		for y := 0; y < len(strs); i++ {
			if string(tmp) != string(strs[y][i]) {
				isOk = false
			}
		}
		if isOk {
			prifix += string(tmp)
		} else {
			break
		}
	}
	return prifix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
