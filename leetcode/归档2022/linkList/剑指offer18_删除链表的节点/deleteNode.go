package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(num []int) *ListNode {
	head := &ListNode{}
	curr := head
	for i := 0; i < len(num); i++ {
		curr.Next = &ListNode{
			Val: num[i],
		}
		curr = curr.Next
	}
	return head
}

func printList(node *ListNode) {
	for node != nil {
		if node.Next == nil {
			fmt.Println(node.Val)
		} else {
			fmt.Println(node.Val, " -> ")
		}
	}
}

func deleteNode(head *ListNode, val int) *ListNode {
	prev := head
	curr := head

	for curr != nil {
		if curr.Val == val {
			if curr == head {
				return head.Next
			} else if curr.Next == nil {
				prev.Next = nil
				return head
			} else {
				prev.Next = curr.Next
				return head
			}
		}
		prev = curr
		curr = curr.Next
	}
	return curr
}
