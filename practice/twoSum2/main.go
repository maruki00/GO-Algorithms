package main

func twoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)

	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}
		if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}
	return []int{}

}

func main() {

}
