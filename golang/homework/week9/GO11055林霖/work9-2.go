package main

import "fmt"

// 单向 node
type SingleNode struct {
	Next  *SingleNode
	Value string
}

func NewSingleNode(value string) *SingleNode {
	return &SingleNode{Value: value}
}

//单向链表
type SingleNodeList struct {
	header *SingleNode
	tail   *SingleNode
}

//append
func (s *SingleNodeList) append(node *SingleNode) *SingleNodeList {
	if s.tail != nil {
		s.tail.Next = node
		s.tail = node
	} else {
		s.header = node
		s.tail = node
	}
	return s
}

//遍历
func (s *SingleNodeList) internodes() {
	node := s.header
	for node != nil {
		fmt.Print(node.Value, " ")
		node = node.Next
	}
	fmt.Println()
}
