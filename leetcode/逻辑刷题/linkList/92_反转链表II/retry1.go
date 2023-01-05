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

func reverseList(head *ListNode) {
	prev := &ListNode{}
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	curr := dummy
	index := 0
	var beforeLeft *ListNode
	var afterRight *ListNode
	var subLinkListHead *ListNode
	var subLinkListTail *ListNode
	for index <= right {
		if index == left-1 {
			beforeLeft = curr
			subLinkListHead = curr.Next
		} else if index == right {
			subLinkListTail = curr
			afterRight = curr.Next
			curr.Next = nil
		}
		index++
		curr = curr.Next
	}
	reverseList(subLinkListHead)
	beforeLeft.Next = subLinkListTail
	subLinkListHead.Next = afterRight
	return dummy.Next
}

func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	printList(head)
	printList(reverseBetween(head, 2, 4))
}
