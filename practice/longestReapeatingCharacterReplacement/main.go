package main

import "fmt"

func characterReplacement(s string, k int) int {
	maxCounter := 0
	l, h := 0, 1
	diffs := 0
	currLitter := s[0]

	for h < len(s) && l < h {

		if currLitter == s[h] || diffs < k {
			if s[h] != currLitter {
				diffs++
			}
			h++
			maxCounter = max(maxCounter, h-l)
			continue
		}

		if s[l] != currLitter {
			diffs--
		}
		maxCounter = max(maxCounter, h-l)

		l++

	}

	return maxCounter

}

func main() {
	s := "AABABBA"
	k := 1
	fmt.Println("result : ", characterReplacement(s, k))
}
