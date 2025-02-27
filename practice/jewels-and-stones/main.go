package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	result := 0
	stonesCount := make(map[rune]int)

	for _, stone := range jewels {
		stonesCount[stone]++
	}

	for _, stone := range stones {
		result += stonesCount[stone]
	}

	return result
}

func main() {
	jewels := "aA"
	stones := "aAAbbbb"

	fmt.Println("result : ", numJewelsInStones(jewels, stones))
}
