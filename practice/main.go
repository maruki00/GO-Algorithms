package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next.Next
	for fast != nil && slow != nil {
		if slow.Val == fast.Val {
			return true
		}

		if slow == nil || fast == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
	return false

}

func main() {

}
