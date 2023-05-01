package main

import (
	"fmt"
)

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

func reversePrint(head *ListNode) []int {
	lists := []int{}
	for head != nil {
		lists = append(lists, head.Val)
		head = head.Next
	}
	ansLists := []int{}
	//for i, j := 0, len(lists)-1; i < len(lists) && j >= 0; i, j = i+1, j-1 {
	//	ansLists[i] = lists[j]
	//}
	for i := len(lists) - 1; i >= 0; i-- {
		ansLists = append(ansLists, lists[i])
	}
	return ansLists
}
