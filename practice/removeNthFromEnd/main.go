package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	lowCount := 0
	heightCount := 0

	lowPointer := head
	heightPointer := head

	for heightPointer != nil {
		for lowCount < heightCount-(n+1) {
			lowPointer = lowPointer.Next
			lowCount++
		}

		heightCount++
		heightPointer = heightPointer.Next
	}

	lowPointer = lowPointer.Next.Next

	return head

}

func main() {
	
}
