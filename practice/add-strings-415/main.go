package main

import (
	"fmt"
	"math/big"
)

func addStrings(num1 string, num2 string) string {
	l1, l2 := len(num1)-1, len(num2)-1
	result := ""
	l := 1
	remain := 0
	tmp := 0
	for l1 >= 0 && l2 >= 0 {
		tmp = int(num1[l1]-'0') + int(num2[l2]-'0')
		result = string((tmp+remain)%10+'0') + result
		remain = (remain + tmp) / 10

		fmt.Println(tmp, int(num1[l1]-'0'), int(num2[l2]-'0'), remain, result)
		l1--
		l2--
		l *= 10

	}
	for l1 >= 0 {
		tmp = int(num1[l1] - '0')
		result = string((tmp+remain)%10+'0') + result
		remain = (remain + tmp) / 10
		l1--
		l *= 10
	}
	for l2 >= 0 {
		tmp = int(num2[l2] - '0')
		result = string((tmp+remain)%10+'0') + result
		remain = (remain + tmp) / 10
		l2--
		l *= 10
	}
	return result
	//return fmt.Sprintf("%s", result)
}

func _addStrings(a string, b string) string {
	i := big.NewInt(0)
	j := big.NewInt(0)
	i.SetString(a, 10)
	j.SetString(b, 10)
	res := i.Add(i, j)
	return fmt.Sprintf("%v", res)
}

func main() {
	fmt.Println(addStrings("3876620623801494171", "6529364523802684779"))
}
