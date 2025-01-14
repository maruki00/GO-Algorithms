package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashMap := make(map[*ListNode]bool, 0)
	first, second := headA, headB
	for first != nil || second != nil {

		if first != nil {
			if _, ok := hashMap[first]; ok {
				return first
			}
			hashMap[first] = true
			first = first.Next
		}
		if second != nil {
			if _, ok := hashMap[second]; ok {
				return second
			}
			hashMap[second] = true
			second = second.Next
		}
	}
	return nil
}

func main() {

}
