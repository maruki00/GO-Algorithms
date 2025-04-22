package main

import "fmt"

type Trie struct {
	root  map[rune]*Trie
	isEnd bool
}

func NewTrie() *Trie {
	return &Trie{root: make(map[rune]*Trie)}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		if node.root[ch] == nil {
			node.root[ch] = &Trie{root: make(map[rune]*Trie)}
		}
		node = node.root[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		if node.root[ch] == nil {
			return false
		}
		node = node.root[ch]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, ch := range prefix {
		if node.root[ch] == nil {
			return false
		}
		node = node.root[ch]
	}
	return true
}

func main() {
	trie := NewTrie()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // true
	fmt.Println(trie.Search("app"))     // false
	fmt.Println(trie.StartsWith("app")) // true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // true
}
