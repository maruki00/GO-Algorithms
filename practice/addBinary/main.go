package main

import (
	"fmt"
	"math/big"
)

func addBinary(a string, b string) string {
	i := big.NewInt(0)
	j := big.NewInt(0)
	i.SetString(a, 2)
	j.SetString(b, 2)
	res := i.Add(i, j)
	return fmt.Sprintf("%b", res)
}

func main() {

	fmt.Println(addBinary("000010100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
		"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}
