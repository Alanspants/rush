package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(nums []int) *ListNode {
	pre := ListNode{}
	head := &pre
	for _, v := range nums {
		head.Next = &ListNode{
			Val: v,
		}
		head = head.Next
	}
	return pre.Next
}

func ShowList(head *ListNode) {
	for head != nil {
		if head.Next == nil {
			fmt.Println(head.Val)
		} else {
			fmt.Print(head.Val, " -> ")
		}
		head = head.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := ListNode{}
	ans := &pre
	overflow := 0
	v1, v2 := 0, 0
	for {
		if l1 == nil && l2 == nil {
			if overflow != 0 {
				ans.Next = &ListNode{
					Val: overflow,
				}
				ans = ans.Next
			}
			break
		}
		if l1 == nil {
			v1 = 0
			v2 = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			v2 = 0
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = l1.Val
			v2 = l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}
		tempSum := v1 + v2 + overflow
		if tempSum >= 10 {
			tempSum -= 10
			overflow = 1
		} else {
			overflow = 0
		}
		ans.Next = &ListNode{
			Val: tempSum,
		}
		ans = ans.Next
	}
	return pre.Next
}

func main() {
	l1 := CreateList([]int{2, 4, 3})
	l2 := CreateList([]int{5, 6, 4})
	ShowList(l1)
	ShowList(l2)
	ans := addTwoNumbers(l1, l2)
	ShowList(ans)
}
