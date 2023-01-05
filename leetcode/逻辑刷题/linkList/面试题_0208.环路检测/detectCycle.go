package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(num []int) *ListNode {
	prev := &ListNode{}
	curr := prev
	for _, n := range num {
		curr.Next = &ListNode{
			Val: n,
		}
		curr = curr.Next
	}
	return prev.Next
}

func printList(node *ListNode) {
	for node != nil {
		if node.Next == nil {
			fmt.Println(node.Val)
		} else {
			fmt.Print(node.Val, " -> ")
		}
		node = node.Next
	}
}

func detectCycle(head *ListNode) *ListNode {

}
