package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	nodes := make(map[*Node]*Node)
	counter := 0
	h := head
	var hh *Node
	for h != nil {

		Nextnode, ok := nodes[h.Next]
		if !ok {
			var newNode *Node
			if Nextnode == nil {
				newNode = nil
			} else {
				newNode = &Node{
					Val:    h.Next.Val,
					Next:   h.Next.Next,
					Random: h.Next.Random,
				}
			}

			nodes[newNode] = newNode

		}

		randomNode, ok := nodes[h.Random]
		if !ok {
			var newNode *Node
			if randomNode == nil {
				newNode = nil
			} else {
				newNode = &Node{
					Val:    h.Random.Val,
					Next:   h.Random.Next,
					Random: h.Random.Random,
				}
			}

			nodes[newNode] = newNode
		}

		node := &Node{
			Val:    h.Val,
			Next:   nodes[h.Next],
			Random: nodes[h.Random],
		}

		nodes[node] = node
		hh = node
		hh = hh.Next
		counter++
		h = h.Next
	}

	return hh
}

func main() {
	next := &Node{Val: 1, Next: nil, Random: nil}
	node := &Node{Val: 10, Next: next, Random: nil}

	fmt.Println("result : ", copyRandomList(node))
}
