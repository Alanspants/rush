package main

import (
	"fmt"
)

// 双向node
type DoubleNode struct {
	Pre   *DoubleNode
	Value string
	Next  *DoubleNode
}

func NewDoubleNode(value string) *DoubleNode {
	return &DoubleNode{Value: value}
}

// 双向链表
type DoubleNodeList struct {
	size   int
	header *DoubleNode
	tail   *DoubleNode
}

// append
func (d *DoubleNodeList) append(node *DoubleNode) *DoubleNodeList {

	if d.tail != nil {
		d.tail.Next = node
		node.Pre = d.tail
		d.tail = node
	} else {
		d.header = node
		d.tail = node
	}
	d.size++
	return d
}

// 遍历元素
func (d *DoubleNodeList) interNodes() {
	node := d.header
	for node != nil {
		fmt.Print(node.Value, " ")
		node = node.Next
	}
	fmt.Println()
}

// 弹出最后一个
func (d *DoubleNodeList) pop() bool {
	if d.size == 1 {
		d.header = nil
		d.tail = nil
		d.size--
		return true
	} else if d.size <= 0 {
		fmt.Println("当前无元素或size异常")
		return false
	} else {
		pre := d.tail.Pre
		pre.Next = nil
		d.tail = pre
		d.size--
		return true
	}
}

// 在下标后面插入元素，下标从1开始，0代表在所有元素最前面
func (d *DoubleNodeList) insert(index int, node *DoubleNode) bool {
	if index < 0 {
		fmt.Println("索引下标不能小于0")
		return false
	}

	if d.size == index {
		d.append(node)
		return true
	} else if d.size < index {
		panic("数组越界")
	}

	if index == 0 {
		h := d.header
		node.Next = h
		h.Pre = node
		d.header = node
		d.size++
		return true
	}

	curIndex := 1
	n := d.header
	for curIndex < index {
		n = n.Next
		curIndex++
	}
	next := n.Next
	n.Next = node
	node.Pre = n
	node.Next = next
	next.Pre = node
	d.size++
	return true

}

// 移除下标的值，下标从1开始
func (d *DoubleNodeList) remove(index int) bool {

	if index <= 0 || index > d.size {
		panic("数组越界或index 小于等于0")
	}

	if index == d.size {
		return d.pop()
	}

	curIndex := 1
	n := d.header
	for curIndex < index {
		n = n.Next
		curIndex++
	}
	pre := n.Pre
	pre.Next = n.Next
	n.Next.Pre = pre
	n.Pre = nil
	d.size--
	return true

}
