package main

const Alphabites = 26

type Node struct {
	children [Alphabites]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{&Node{}}
}

func (obj *Trie) insert(w string) {
	wordLenght := len(w)
	currebtNode := obj.root
	for i := 0; i < wordLenght; i++ {
		charIndex := w[i] - 'a'
		if currebtNode.children[charIndex] == nil {
			currebtNode.children[charIndex] = &Node{}
		}
		currebtNode = currebtNode.children[charIndex]
	}
	currebtNode.isEnd = true
}

func (obj *Trie) search(w string) bool {
	wordLenght := len(w)
	currebtNode := obj.root
	for i := 0; i < wordLenght; i++ {
		charIndex := w[i] - 'a'
		if currebtNode.children[charIndex] == nil {
			return false
		}
		currebtNode = currebtNode.children[charIndex]
	}
	if currebtNode.isEnd == true {
		return true
	}
	return false
}

func main() {

}
