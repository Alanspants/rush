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
	return head.Next
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		if curr.Next == nil {
			fmt.Println(curr.Val)
		} else {
			fmt.Print(curr.Val, " -> ")
		}
		curr = curr.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	nodeMap := map[int]int{}

	curr := head

	for curr != nil {
		if _, ok := nodeMap[curr.Val]; ok {
			nodeMap[curr.Val] += 1
		} else {
			nodeMap[curr.Val] = 1
		}
		curr = curr.Next
	}

	prev := head
	curr = head

	for curr != nil {
		if nodeMap[curr.Val] != 1 {
			if curr == head {
				if curr.Next != nil {
					head = head.Next
					curr = curr.Next
					prev = nil
				} else {
					return nil
				}
			} else {
				if curr.Next != nil {
					curr = curr.Next
					prev.Next = curr
				} else {
					prev.Next = nil
					return head
				}
			}
		} else {
			prev = curr
			curr = curr.Next
		}
	}
	return head
}

func main() {
	nums := []int{1, 2, 2}
	node := createList(nums)
	printList(node)
	ans := deleteDuplicates(node)
	printList(ans)
}
