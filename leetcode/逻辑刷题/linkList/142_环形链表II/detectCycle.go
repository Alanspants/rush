package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	hashmap := map[*ListNode]bool{}

	curr := head
	for curr != nil {
		if _, ok := hashmap[curr]; ok {
			return curr
		} else {
			hashmap[curr] = true
		}
		curr = curr.Next
	}
	return nil
}
