package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	sen := &ListNode{0, head}
	prev, cur := sen, head
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}

	return sen.Next

}

func main() {

}
