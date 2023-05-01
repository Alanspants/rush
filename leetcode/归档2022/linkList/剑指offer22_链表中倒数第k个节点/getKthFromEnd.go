package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	quick := head
	slow := head

	for i := 0; i < k; i++ {
		quick = quick.Next
	}

	for quick != nil {
		quick = quick.Next
		slow = slow.Next
	}

	return slow
}
