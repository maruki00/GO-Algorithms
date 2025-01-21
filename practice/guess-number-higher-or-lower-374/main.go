package main

import "fmt"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 */
func guess(n int) int {
	gess := 6
	if n == gess {
		return 0
	}
	if n > gess {
		return -1
	}
	return 1
}

func guessNumber(n int) int {

	start, end := 0, n
	for start < end {
		middle := int((end + start) / 2)
		guessed := guess(middle)
		if guessed == 0 {
			return n
		}
		if guessed == -1 {
			start = middle + 1
			continue
		}
		end = middle - 1
	}
	return n
}

func main() {
	fmt.Println(guessNumber(8))
}
