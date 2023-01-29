package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	quick := head.Next
	slow := head
	prev := slow

	for quick != nil && slow != nil {
		if quick == slow {
			return prev
		}
		if quick.Next == nil || quick.Next.Next == nil {
			return nil
		}
		quick = quick.Next.Next
		prev = slow
		slow = slow.Next
	}
	return nil
}
