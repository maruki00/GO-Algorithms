package main

import (
	"fmt"
	"strings"
	"unicode"
)

const MAXINT = 2147483647
const MININT = -2147483648

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	result := 0

	if s == "" {
		return result
	}

	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	for _, char := range s {
		if !unicode.IsDigit(char) {
			break
		}

		result = ((result * 10) + int(char-'0'))

		if sign*result > MAXINT {
			return MAXINT
		} else if sign*result < MININT {
			return MININT
		}

	}
	return result * sign
}

func main() {
	fmt.Println(myAtoi("-91283472332"))
}
