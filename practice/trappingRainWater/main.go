package main

import "fmt"

func trap(height []int) int {

	fmt.Println("len : ", len(height))
	l := 0
	r := 1
	resu := 0
	for l < len(height)-1 && r < len(height) {
		// if l-r < 0 {
		// 	l++
		// 	r++
		// 	continue
		// }

		for r < len(height) && height[r] <= height[l] {
			r++
		}
		r--
		minHieght := min(height[r], height[l])
		for l < r-1 && l < len(height)-1 && r < len(height) {
			resu += minHieght - height[l]
			l++
		}
		l = r
		r++

	}

	return resu

}

func main() {
	items := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("result : ", trap(items))
}
