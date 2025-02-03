package main

func Sort(nums []int) {

	for i := 0; i < len(nums); i++ {
		min_index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[min_index] {
				min_index = j
			}
		}
	}

}

func main() {

}
