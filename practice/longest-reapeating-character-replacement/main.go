package main

import "fmt"

func characterReplacement(s string, k int) int {

	maxCounter := 0
	chars := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if chars[s[i]] {
			continue
		}
		chars[s[i]] = true

		l, r, diff := 0, 0, 0

		for r < len(s) && l < len(s) {
			if s[r] == s[i] {
				r++
			} else if diff < k {
				diff++
				r++
			}else if 
		}
	}

	return maxCounter
}

func characterReplacement2(s string, k int) int {
	res := 0
	for i := 0; i < 26; i++ {
		cur_k := k
		max_len := 0
		cur_char := 'A' + byte(i)
		l := 0
		r := 0
		for r < len(s) && l < len(s) {
			if s[r] == cur_char {
				r++
			} else if cur_k > 0 {
				r++
				cur_k--
			} else if s[l] == cur_char {
				l++
			} else {
				l++
				cur_k++
			}
			max_len = max(max_len, r-l)
		}
		res = max(res, max_len)
	}
	return res
}

func main() {
	s := "ABAA"
	k := 0
	fmt.Println("result : ", characterReplacement(s, k))
}
