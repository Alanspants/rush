package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(nums []int) *ListNode {
	prev := &ListNode{}
	curr := prev
	for _, num := range nums {
		curr.Next = &ListNode{
			Val: num,
		}
		curr = curr.Next
	}
	return prev.Next
}

func printList(head *ListNode) {
	for head != nil {
		if head.Next == nil {
			fmt.Println(head.Val)
		} else {
			fmt.Print(head.Val, " -> ")
		}
		head = head.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	listNode := createList([]int{1, 2, 3, 4, 5})
	printList(listNode)
	listNode = reverseList(listNode)
	printList(listNode)
}
