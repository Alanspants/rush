package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tempSum := 0
	overFlow := 0
	header := &ListNode{}
	current := header
	for l1 != nil && l2 != nil {
		if l1 == nil {
			tempSum = 0 + l2.Val + overFlow
			l2 = l2.Next
		} else if l2 == nil {
			tempSum = l1.Val + 0 + overFlow
			l1 = l1.Next
		} else {
			tempSum = l1.Val + l2.Val + overFlow
			l1 = l1.Next
			l2 = l2.Next
		}
		if tempSum >= 10 {
			overFlow = 1
			tempSum -= 10
		} else {
			overFlow = 0
		}
		current.Next = &ListNode{tempSum, nil}
		current = current.Next
	}
	if overFlow == 1 {
		current.Next = &ListNode{1, nil}
		current = current.Next
	}

	return header.Next
}

func main() {
	node1 := CreateNode(2)
	node2 := CreateNode(4)
	node3 := CreateNode(3)
	node1.Next = node2
	node2.Next = node3

	node11 := CreateNode(5)
	node12 := CreateNode(6)
	node13 := CreateNode(4)
	node11.Next = node12
	node12.Next = node13
}
