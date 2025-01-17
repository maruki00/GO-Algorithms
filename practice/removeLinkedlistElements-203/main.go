package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	tmp := head

	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		}
	}
	return head
}

func main() {

}
