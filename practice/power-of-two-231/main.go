package main

import "math"

func isPowerOfTwo(n int) bool {
	res := n & (n - 1)
	return res == 0 && n != 0
}
func main() {

}
