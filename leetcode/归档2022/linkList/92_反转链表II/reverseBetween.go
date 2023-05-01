package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createNode(nums []int) *ListNode {
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

func printNode(node *ListNode) {
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
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}
	index := 0
	pre := dummyNode

	for index < left-1 {
		pre = pre.Next
		index++
	}
	leftNode := pre.Next
	rightNode := pre
	for index < right {
		rightNode = rightNode.Next
		index++
	}
	succ := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil

	reverseList(leftNode)

	pre.Next = rightNode
	leftNode.Next = succ

	return dummyNode.Next
}

func main() {
	listNode := createNode([]int{1, 2, 3, 4, 5})
	printNode(listNode)
	listNode = reverseBetween(listNode, 2, 4)
	printNode(listNode)
}
