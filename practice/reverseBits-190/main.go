package main

func reverseBits(num uint32) uint32 {
	result := uint32(0)
	for num > 0 {
		result += uint32((uint32(result) * 10) + (num % 10))
		num /= 10
	}
	return result
}

func main() {}
