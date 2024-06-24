package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

const MAXINT = 2147483647
const MININT = -2147483648

func myAtoi_(s string) int {
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
		if char < '0' || char > '9' {
			break
		}

		result = ((result * 10) + int(char-'0'))

		if sign*result > MAXINT {
			return MAXINT
		} else if sign*result < MININT {
			return MININT
		}

	}
	return (result) * sign
}
func myAtoi(s string) int {
	sign := 1
	res := 0
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return res
	}

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			break
		}
		dig := int(ch - '0')
		res = res*10 + dig

		if sign*res < math.MinInt32 {
			return math.MinInt32
		} else if sign*res > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return sign * res
}

func main() {
	fmt.Println(myAtoi("-91283472332"))
}
