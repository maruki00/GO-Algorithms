package main

import "fmt"

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
	return curr
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println("result : ", compress(chars))

}
