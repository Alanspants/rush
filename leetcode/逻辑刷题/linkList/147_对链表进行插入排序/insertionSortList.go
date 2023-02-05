package main

import (
	"fmt"
)

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

func insertionSortList(head *ListNode) *ListNode {
	curr := head
	prev := head

	if curr.Next == nil {
		return head
	} else {
		curr = curr.Next
		prev = head
	}

	for curr != nil {
		tempCurr := head
		tempPrev := head
		for tempCurr != curr {
			if curr.Val <= tempCurr.Val {
				if tempCurr == head {
					if curr.Next != nil {
						dummyPrev := prev
						dummyCurr := curr.Next

						prev.Next = curr.Next
						head = curr
						curr.Next = tempCurr

						prev = dummyPrev
						curr = dummyCurr
						break
					} else {
						dummyPrev := prev
						dummyCurr := curr
						dummyCurr = nil

						prev.Next = nil
						head = curr
						curr.Next = tempCurr

						prev = dummyPrev
						curr = dummyCurr
						break
					}
				} else {
					if curr.Next != nil {
						dummyPrev := prev
						dummyCurr := curr.Next

						prev.Next = curr.Next
						tempPrev.Next = curr
						curr.Next = tempCurr

						prev = dummyPrev
						curr = dummyCurr
						break
					} else {
						dummyPrev := prev
						dummyCurr := curr
						dummyCurr = nil

						prev.Next = nil
						tempPrev.Next = curr
						curr.Next = tempCurr

						prev = dummyPrev
						curr = dummyCurr
						break
					}
				}
			}
			tempPrev = tempCurr
			tempCurr = tempCurr.Next
		}
		if tempCurr == curr {
			prev = curr
			curr = curr.Next
		}
	}
	return head
}

func main() {
	nums := []int{-1, 5, 3, 4, 0}
	head := createList(nums)
	printList(head)
	ans := insertionSortList(head)
	printList(ans)
}
