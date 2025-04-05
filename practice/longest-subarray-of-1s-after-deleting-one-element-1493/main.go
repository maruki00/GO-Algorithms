package main

func longestSubarray(nums []int) int {
	i, j := 0, 0
	numOfZeros := 0
	maxOnes := 0
	for j < len(nums) {
		for nums[i] != 0 {
			i++
		}
		j = i
		for nums[j] == 1 {
			j++
		}
		maxOnes = max(maxOnes, j-i)
		i = j
	}
	return maxOnes
}

func main() {

}
