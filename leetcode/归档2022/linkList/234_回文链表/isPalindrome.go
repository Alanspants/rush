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

func isPalindrome(head *ListNode) bool {
	origin := []int{}
	curr := head
	for curr != nil {
		origin = append(origin, curr.Val)
		curr = curr.Next
	}
	afterReverse := reverseList(head)
	for i := 0; i < len(origin); i++ {
		if origin[i] != afterReverse.Val {
			return false
		}
		afterReverse = afterReverse.Next
	}

	return true
}

func main() {
	num := []int{1, 2, 3, 4, 5}
	head := createList(num)
	printList(head)
	head = reverseList(head)
	printList(head)
}
