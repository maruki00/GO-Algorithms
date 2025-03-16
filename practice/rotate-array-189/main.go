package main

func rotate(nums []int, k int) {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[(i+k)%n] = nums[i]
	}
	copy(nums, result)
}
func rotate_optimized(nums []int, k int) {
	n := len(nums)
	i, j := 0, len(nums)-1
	midd := k
	if k > n {
		midd = k % n
	}
	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	i, j = 0, midd-1
	for i <= j && j < n {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	i, j = midd, n-1
	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {

}
