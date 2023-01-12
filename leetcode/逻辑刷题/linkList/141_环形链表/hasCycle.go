package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func hasCycle(head *ListNode) bool {
//	hashmap := map[*ListNode]bool{}
//	for head != nil {
//		if _, OK := hashmap[head]; OK {
//			return true
//		}
//		hashmap[head] = true
//		head = head.Next
//	}
//	return false
//}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	quickNode := head.Next
	slowNode := head
	for quickNode != nil && slowNode != nil {
		if quickNode == slowNode {
			return true
		}
		if quickNode == nil || quickNode.Next == nil {
			return false
		}
		quickNode = quickNode.Next.Next
		slowNode = slowNode.Next
	}
	return false
}
