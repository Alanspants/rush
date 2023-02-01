package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//curr := head
	//length := 0
	//for curr != nil {
	//	curr = curr.Next
	//	length++
	//}
	//if n == length {
	//	if
	//}
	quick := head
	slow := head
	prev := slow
	for i := 0; i < n; i++ {
		quick = quick.Next
	}
	for quick != nil {
		quick = quick.Next
		prev = slow
		slow = slow.Next
	}

	if slow == head {
		return head.Next
	} else if slow.Next == nil {
		prev.Next = nil
	} else {
		prev.Next = slow.Next
	}

	return head
}
