package main

import "unicode"

type Stack struct {
	items []byte
}

func NewStack() *Stack {
	return &Stack{
		items: make([]byte, 0),
	}
}

func (s *Stack) Push(char byte) {
	s.items = append(s.items, char)
}
func (s *Stack) Pop() byte {
	lst := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return lst
}

func reverseVowels_(s string) string {
	voiles := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,

		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		if _, ok := voiles[s[i]]; ok {
			stack.Push(s[i])
		}
	}
	result := ""
	for i := 0; i < len(s); i++ {
		if _, ok := voiles[s[i]]; ok {
			tmp := stack.Pop()
			result += string(tmp)
			continue
		}
		result += string(s[i])
	}
	return result

}

func isVoile(v byte) bool {
	char := unicode.ToLower(rune(v))
	return char == 'a' ||
		char == 'e' ||
		char == 'i' ||
		char == 'o' ||
		char == 'u'
}
func reverseVowels(s string) string {
	// voiles := map[byte]bool{
	// 	'a': true,
	// 	'e': true,
	// 	'i': true,
	// 	'o': true,
	// 	'u': true,

	// 	'A': true,
	// 	'E': true,
	// 	'I': true,
	// 	'O': true,
	// 	'U': true,
	// }
	out := []byte(s)
	start, end := 0, len(s)-1
	for start < end {

		if !isVoile(s[start]) {
			start++
		}
		if !isVoile(s[end]) {
			end--
		}
		if isVoile(s[end]) && isVoile(s[start]) {
			out[start], out[end] = out[end], out[start]
			start++
			end--
		}
	}
	return string(out)
}

func main() {

}
