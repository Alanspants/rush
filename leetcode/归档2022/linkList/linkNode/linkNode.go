package linkNode

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func CreateList(nums []int) LinkNode {
	pre := LinkNode{
		Val: 0,
	}

	cur := &pre

	for _, v := range nums {
		cur.Next = &LinkNode{
			Val: v,
		}
		cur = cur.Next
	}

	return *pre.Next
}

func PrintList(head *LinkNode) {
	for {
		if head.Next != nil {
			fmt.Print(head.Val, " -> ")
			head = head.Next
		} else {
			fmt.Println(head.Val)
			break
		}
	}
}
