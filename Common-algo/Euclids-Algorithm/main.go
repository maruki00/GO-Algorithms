package main

import (
	"fmt"
	"math/big"
)

func gcd(a, b *big.Int) *big.Int {
	index := 0
	temp := new(big.Int)

	for b.Cmp(big.NewInt(0)) != 0 {
		temp.Mod(a, b)
		a.Set(b)
		b.Set(temp)
		index++
	}
	println("times : ", index)
	return a
}

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("239837405689734523452345234985723049854", 10)
	b.SetString("4334958763085670349758603498756034856703485670349857089703845760384576034857608457603845756", 10)

	result := gcd(a, b)
	fmt.Printf("The GCD is: %s\n", result.String())
}
