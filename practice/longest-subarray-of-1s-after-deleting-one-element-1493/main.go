package main

func longestSubarray(nums []int) int {
	j := 0
	numOfZeros := 0
	maxOnes := 0
	for i:=0;i<len(nums);i++ {
		for nums[j] != 0 {
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
