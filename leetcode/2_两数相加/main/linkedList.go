package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Header *ListNode
}

func CreateNode(v int) *ListNode {
	return &ListNode{v, nil}
}

func CreateLinkedList(v int) *LinkedList {
	header := CreateNode(v)
	return &LinkedList{header}
}

func (llist *LinkedList) AddNode(v int) {
	node := CreateNode(v)
	temp := llist.Header
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = node
}
