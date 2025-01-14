package main

import "fmt"

func reverseBits(num uint32) uint32 {
	result := uint32(0)
	for num > 0 {
		fmt.Println(num, result)
		result = uint32((result * 10) + (num % 10))
		num /= 10
	}
	return result
}

func main() {

	y := 43261596
	fmt.Println(reverseBits(uint32(y)))
}
