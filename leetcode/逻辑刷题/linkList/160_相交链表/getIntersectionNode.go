package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode
//	curr := head
//	for curr != nil {
//		next := curr.Next
//		curr.Next = prev
//		prev = curr
//		curr = next
//	}
//	return prev
//}
//
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	headAReverse := reverseList(headA)
//	headBReverse := reverseList(headB)
//	currentA := headAReverse
//	currentB := headBReverse
//	for currentA != nil && currentB != nil {
//		if currentA == currentB {
//			headA = reverseList(headAReverse)
//			headB = reverseList(headBReverse)
//			return currentA
//		} else {
//			currentA = currentA.Next
//			currentB = currentB.Next
//		}
//	}
//	return nil
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currentA, currentB := headA, headB

	for currentA != currentB {
		if currentA != nil {
			currentA = currentA.Next
		} else {
			currentA = headB
		}
		if currentB != nil {
			currentB = currentB.Next
		} else {
			currentB = headA
		}
	}

	if currentA != nil {
		return currentA
	} else {
		return nil
	}
}
