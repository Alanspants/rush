package main

import (
	"errors"
	"fmt"
)

type node struct {
	prev  *node
	next  *node
	value interface{}
}
type LinkedList struct {
	head *node
	tail *node
	len  int
}
type Iternodes struct {
	data  *LinkedList
	index int
}

func main() {
	link := LinkedList{}
	fmt.Println(link)
	link.Append(1)
	link.Append(2)
	link.Append(3)
	link.Append(4)
	link.Append(5)
	link.Insert(6, 6)
	link.Append(7)
	link.Append(8)
	link.Insert(7, 100)
	link.Remove(3)
	for it := link.Iternodes(); it.HasNext(); {
		fmt.Println(it.Next())

	}
	fmt.Println(link)

}

//func NewLinkList(name string, age int) *LinkedList {
//
//	return &LinkedList{name: name, age: age}
//}
func (this *LinkedList) Append(value any) {
	node := &node{value: value}
	if this.len == 0 {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		node.prev = this.tail
		this.tail = node
	}
	this.len++
}
func (this *LinkedList) Iternodes() *Iternodes {

	return &Iternodes{data: this, index: 0}
}
func (this *LinkedList) Pop() error {
	if this.head == nil {
		return errors.New("this link is empty")
	}
	newtail := this.tail.prev
	this.tail.prev.next = nil
	this.tail = newtail
	this.len--
	return nil
}
func (this *LinkedList) Insert(i int, value any) error {
	newNode := &node{value: value}
	lnode, err := this.Findbyid(i)
	fmt.Printf("node.value: %v\n", lnode.value)
	if err != nil {
		return err
	}
	if i <= 1 {
		this.head = newNode
	}
	newNode.next = lnode
	newNode.prev = lnode.prev
	lnode.prev.next = newNode
	lnode.prev = newNode
	this.len++
	return nil
}
func (this *LinkedList) Remove(index int) error {
	if index > this.len || index < 0 {
		return errors.New("index out of range")
	}
	node, err := this.Findbyid(index)
	if err != nil {
		return err
	}
	if index == 0 {
		this.head = node.next
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	return nil

}
func (this *LinkedList) Findbyid(index int) (*node, error) {
	if index <= 0 { //索引小于=0,返回头
		return this.head, nil
	} else if index >= this.len { //索引大于或登录list长度,返回尾
		return this.tail, nil
	} else {
		node := this.head
		for i := 1; i < index; i++ {
			node = node.next

		}
		return node, nil
	}

}
func (this *Iternodes) HasNext() bool {
	return this.index < this.data.len-1
}

var xnode *node

func (this *Iternodes) Next() (v any) {
	if this.index == 0 {
		xnode = this.data.head
	}
	v = xnode.value
	if xnode.next == nil {
		this.index++
		return v
	}
	xnode = xnode.next

	this.index++
	return v
}
