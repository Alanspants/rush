//package main

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
	curr := head
	nodes := []*ListNode{}
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	//for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
	//	nodes[i].Next = nodes[j]
	//	nodes[j].Next = nodes[i+1]
	//}
	i := 0
	j := len(nodes) - 1
	for i <= j {
		nodes[i].Next = nodes[j]
		if i == j {

			break
		}
		i++
		nodes[j].Next = nodes[i]
		j--
	}

	nodes[i].Next = nil
}

func main() {
	listNode := createList([]int{1, 2, 3, 4})
	printList(listNode)
	reorderList(listNode)
	printList(listNode)
}
