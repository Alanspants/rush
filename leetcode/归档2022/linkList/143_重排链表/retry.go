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

func reorderList(head *ListNode) {
	lists := []*ListNode{}
	curr := head
	for curr != nil {
		lists = append(lists, curr)
		curr = curr.Next
	}

	for i, j := 0, len(lists)-1; i <= j; i, j = i+1, j-1 {
		lists[i].Next = lists[j]
		if i+1 <= j-1 {
			lists[j].Next = lists[i+1]
		} else {
			lists[j].Next = nil
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	head := createList(nums)
	printList(head)
	reorderList(head)
	printList(head)
}
