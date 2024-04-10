package main

import "fmt"

func main() {
	var items []int
	items = append(items, 10)
	items = append(items, 15)
	items = append(items, 20)
	items = append(items, 25)
	fmt.Println("Items: ", items)
}
