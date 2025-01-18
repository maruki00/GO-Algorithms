package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	nodes := make([]*ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	first, last := 0, len(nodes)-1
	for first <= last {
		if nodes[first] != nodes[last] {
			return false
		}
		first++
		last--
	}
	return true
}

func main() {

}
