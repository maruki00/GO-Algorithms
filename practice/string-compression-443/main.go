package main

import (
	"fmt"
)

func compress(chars []byte) int {
	ln := len(chars)
	if ln == 0 {
		return 0
	}
	pos, curr, next := 0, 0, 0
	for next < ln {
		if chars[curr] == chars[next] {
			next++
			continue
		}
		chars[pos] = chars[curr]
		count := next - curr
		if count == 1 {
			pos++
			curr = next
			continue
		}
		bts := []byte(fmt.Sprintf("%d", count)) // 123 = '1', '2', '3'
		for i, bt := range bts {
			chars[pos+1+i] = bt
		}
		pos += len(bts) + 1
		curr = next
	}
	if next-curr >= 1 {
		chars[pos] = chars[curr]
		if next-curr > 1 {
			bts := []byte(fmt.Sprintf("%d", next-curr)) // 123 = '1', '2', '3'
			for i, bt := range bts {
				chars[pos+1+i] = bt
			}
			pos += len(bts) + 1
		} else {
			pos++
		}

	}
	return pos
}

func main() {
	//chars := []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	chars := []byte{'a', 'b', 'c'}
	fmt.Println("result : ", compress(chars))
	for i, ch := range chars {
		fmt.Println(i, string(ch))
	}

}
