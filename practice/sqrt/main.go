package main

import "fmt"

func mySqrt1(x int) int {
	if x < 2 {
		return x
	}
	y := float64(x)
	var sqrt float64 = float64((y + (float64(x) / float64(y))) / 2)

	for float64(y-sqrt) >= 0.00001 {
		y = sqrt
		sqrt = float64((y + (float64(x) / float64(y))) / 2)

	}
	return int(sqrt)
}

func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		mid := (l + r) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

func main() {
	fmt.Println(mySqrt(8))
}
