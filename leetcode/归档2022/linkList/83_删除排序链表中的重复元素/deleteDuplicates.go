package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(nums []int) *ListNode {
	head := &ListNode{}
	curr := head
	for i := 0; i < len(nums); i++ {
		curr.Next = &ListNode{
			Val: nums[i],
		}
	}
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	slow := head
	quick := head
	if slow != nil && slow.Next != nil {
		quick = slow.Next
	} else {
		return head
	}

	for quick != nil {
		if slow.Val == quick.Val {
			if quick.Next != nil {
				slow.Next = quick.Next
				quick = slow.Next
			} else {
				slow.Next = nil
				return head
			}
		} else {
			slow = quick
			quick = quick.Next
		}
	}
	return head
}
