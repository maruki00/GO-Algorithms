package main

import "fmt"

func letterCombinations(digits string) []string {
	result := []string{}
	symboles := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "jhi",
		'5': "gkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var recur func(syms string, tmp string)
	recur = func(syms string, tmp string) {
		if len(syms) == 0 {
			return
		}
		for _, j := range symboles[syms[0]] {
			tmp += string(j)
			recur(syms[1:], tmp)
			result = append(result, tmp)
			fmt.Println(tmp)
		}
	}
	recur(digits, "")
	return result
}

func main() {
	fmt.Println("result : ", letterCombinations("23"))
}
