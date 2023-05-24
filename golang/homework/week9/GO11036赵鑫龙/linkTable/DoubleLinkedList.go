package linkTable

import (
	"errors"
	"fmt"
	"strconv"
)

type DLinkLister interface {
	Append(value int) *DLinkList
	Remove(index int) error
	Insert(index, value int) error
	Pop() error
	IterNodes(b bool) []*DNode
	Len() int
}
type DNode struct {
	value int
	next  *DNode
	front *DNode
}

type DLinkList struct {
	len  int
	head *DNode
	tail *DNode
}

func (n *DNode) String() string {
	var next string
	var front string
	if n.next == nil {
		next = "null"
	} else {
		next = strconv.Itoa(n.next.value)
	}
	if n.front == nil {
		front = "null"
	} else {
		front = strconv.Itoa(n.front.value)
	}

	return fmt.Sprintf("%s<==%d==>%s", front, n.value, next)
}
func (dl *DLinkList) Len() int {
	return dl.len
}

func (dl *DLinkList) Append(value int) *DLinkList {
	node := new(DNode)
	node.value = value
	//元素为0，头尾都是nil，但凡有一个元素，头尾都不会是nil
	if dl.head == nil {
		//头部和尾部是同一个值
		dl.head = node
		dl.tail = node
	} else {
		dl.tail.next = node  //最后一个值的下一个值变为node
		node.front = dl.tail //最后一个值变为前一个值
		dl.tail = node       //更新尾巴
	}
	dl.len++
	return dl
}
func (dl *DLinkList) Pop() error {
	//链表元素为空时抛出异常
	if dl.len == 0 {
		return errors.New("errors message: 链表为空")
	}
	if dl.head == dl.tail {
		dl.head = nil
		dl.tail = nil
	} else {
		tail := dl.tail     //记录最后一个值
		front := tail.front //记录最后一个值的前一个值
		front.next = nil
		dl.tail = front //将尾部指针移动到前一个值
	}
	dl.len--
	return nil
}
func (dl *DLinkList) Insert(index, value int) error {
	if index < 0 {
		return errors.New("无效值")
	}
	if index >= dl.len {
		dl.Append(value)
		return nil
	}

	var current *DNode
	for i, v := range dl.IterNodes(false) {
		if i == index {
			current = v
			break
		}
	}
	node := new(DNode)
	node.value = value
	if current.front == nil {
		//开头插入
		dl.head = node
	} else {
		//中间插入
		front := current.front
		front.next = node
		node.next = front
	}
	current.front = node
	node.next = current
	dl.len++
	return nil
}

func (dl *DLinkList) Remove(index int) error {
	if dl.tail == nil {
		//空链表
		return errors.New("空链表")
	}
	if index < 0 {
		return errors.New("无效值")
	}
	if index >= dl.len {
		return errors.New("index 超界")
	}
	var current *DNode
	p := dl.head
	for i := 0; i < dl.len; i++ {
		if i == index {
			current = p
			break
		}
		p = p.next
	}
	front := current.front
	next := current.next
	if dl.head == dl.tail {
		dl.head = nil
		dl.tail = nil
	} else if dl.head == current {
		dl.head = next
		next.front = nil
	} else if dl.tail == current {
		front.next = nil
		dl.tail = front
	} else {
		front.next = next
		next.front = front
	}
	dl.len--
	return nil
}
func (dl *DLinkList) IterNodes(b bool) []*DNode {
	var p *DNode
	m := make([]*DNode, 0, dl.len)
	if b == true {
		p = dl.head
	} else {
		p = dl.tail
	}
	for p != nil {
		m = append(m, p)
		if b == true {
			p = p.next
		} else {
			p = p.front
		}
	}
	return m
}
