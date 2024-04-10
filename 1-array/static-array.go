package main

import "fmt"

func main() {
	var items [4]int
	items[0] = 10
	items[1] = 15
	items[2] = 20
	items[3] = 25

	fmt.Println("Items: ", items[:3])
}
