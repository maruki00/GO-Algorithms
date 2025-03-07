package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{}
	symboles := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var recur func(digits string, tmp string)
	recur = func(digits string, tmp string) {
		if len(digits) == 0 {
			result = append(result, tmp)
			return
		}
		for _, j := range symboles[digits[0]] {
			recur(digits[1:], tmp+string(j))
		}
	}
	recur(digits, "")
	return result
}

func main() {
	fmt.Println("result : ", letterCombinations("23"))
}
