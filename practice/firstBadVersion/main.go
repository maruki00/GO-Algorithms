package main

import "math/rand"

func isBadVersion(version int) bool {
	val := rand.Intn(10)

	return val == version
}

func firstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {

}
