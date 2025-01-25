package main

import "fmt"

func intToRoman(num int) string {
	mp := map[int]byte{
		0: '',
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}

	result := ""
	fmt.Println(mp)

	return result
}

func main() {

}
