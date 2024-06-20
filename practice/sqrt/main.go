package main

import "fmt"

func mySqrt(x int) int {
	sqrt := x / 2
	tmp := 0

	for tmp != sqrt {
		tmp = sqrt
		sqrt = (x/tmp + tmp) / 2
	}
	return sqrt
}

func main() {
	fmt.Println(mySqrt(144))
}
