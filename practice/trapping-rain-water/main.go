package main

import "fmt"

func trap(h []int) int {
	res, l, r, lm, rm := 0, 0, len(h)-1, 0, 0
	for l <= r {
		if lm <= rm {
			if h[l] > min(lm, rm) {
				lm = h[l]
			}
			res += max(min(lm, rm)-h[l], 0)
			l++
			continue
		}
		if h[r] > min(lm, rm) {
			rm = h[r]
		}
		res += max(min(lm, rm)-h[r], 0)
		r--

	}
	return res
}
func main() {
	items := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("result : ", trap(items))
}
