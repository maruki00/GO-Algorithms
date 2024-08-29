package main

import "fmt"

func isPerfectSquare(num int) bool {
	l, h := 0, num/2

	for l <= h {
		middle := (l + h) / 2
		sqrt := middle * middle
		if sqrt == num {
			return true
		}

		if sqrt < num {
			l = middle + 1
		} else {
			h = middle - 1
		}
	}

	return false
}

func main() {
	fmt.Println("result :  ", isPerfectSquare(4))
}
