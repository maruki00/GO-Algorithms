package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 + num&1
		num >>= 1
	}
	return res
}

func main() {

	y := 43261596
	fmt.Println(reverseBits(uint32(y)))
}
