package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(nums []int) *ListNode {
	pre := &ListNode{}
	curr := pre
	for _, v := range nums {
		curr.Next = &ListNode{
			Val: v,
		}
		curr = curr.Next
	}
	return pre.Next
}

func printList(head *ListNode) {
	for head != nil {
		if head.Next == nil {
			fmt.Println(head.Val)
			break
		} else {
			fmt.Print(head.Val, " -> ")
			head = head.Next
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp1 := []int{}
	for l1 != nil {
		temp1 = append(temp1, l1.Val)
		l1 = l1.Next
	}
	s1 := []int{}
	for i := range temp1 {
		s1 = append(s1, temp1[len(temp1)-1-i])
	}

	temp2 := []int{}
	for l2 != nil {
		temp2 = append(temp2, l2.Val)
		l2 = l2.Next
	}
	s2 := []int{}
	for i := range temp2 {
		s2 = append(s2, temp2[len(temp2)-1-i])
	}

	pre := &ListNode{}
	ans := pre
	overflow := 0
	index1 := 0
	index2 := 0
	i, j := 0, 0
	sum := 0
	for {
		if index1 == len(s1) && index2 == len(s2) {
			if overflow > 0 {
				ans.Next = &ListNode{
					Val: overflow,
				}
				ans = ans.Next
			}
			break
		}
		if index1 == len(s1) {
			i = 0
			j = s2[index2]
			index2++
		} else if index2 == len(s2) {
			j = 0
			i = s1[index1]
			index1++
		} else {
			i = s1[index1]
			j = s2[index2]
			index1++
			index2++
		}
		if i+j+overflow >= 10 {
			sum = i + j + overflow - 10
			overflow = 1
		} else {
			sum = i + j + overflow
			overflow = 0
		}
		ans.Next = &ListNode{
			Val: sum,
		}
		ans = ans.Next
	}

	temp := []int{}
	pre = pre.Next
	for pre != nil {
		temp = append(temp, pre.Val)
		pre = pre.Next
	}
	preAns := &ListNode{}
	Ans := preAns
	for i := range temp {
		Ans.Next = &ListNode{
			Val: temp[len(temp)-1-i],
		}
		Ans = Ans.Next
	}
	return preAns.Next
}

func main() {
	l1 := createList([]int{7, 2, 4, 3})
	l2 := createList([]int{5, 6, 4})
	printList(l1)
	printList(l2)
	printList(addTwoNumbers(l1, l2))
}
