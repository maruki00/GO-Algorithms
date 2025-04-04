package main

import "fmt"

func compress(chars []byte) int {
	ln := len(chars)
	if ln == 0 {
		return 0
	}
	result := 0
	curr := chars[0]
	count := 0
	for i := 1; i < ln; i++ {
		for curr == chars[i] {
			count++
			continue
		}
		result += 2
		if i < ln {
			break
		}
		curr = chars[i]
		count = 1
	}

	return result
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println("result : ", compress(chars))

}
