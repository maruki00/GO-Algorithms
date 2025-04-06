package main

func uniqueOccurrences(arr []int) bool {
	occures := make(map[int]int)
	for _, num := range arr {
		occures[num]++
	}
	for _, v := range occures {
		if v == 1 {
			return true
		}
	}
	return false
}

func main() {

}
