package main

import "fmt"

func intToRoman(num int) string {
	mp := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	result := ""
	for num > 0 {
		for nu, sumbol := range mp {
			if num-nu >= 0 {
				result += sumbol
				num -= nu
				continue
			}
		}
	}

	return result
}

func main() {
	in := 3749
	fmt.Println("result : ", intToRoman(in))
}
