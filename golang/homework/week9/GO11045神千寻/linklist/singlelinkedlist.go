package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item int
	next *Node
}
type SingleList struct {
	length int
	head   *Node
	tail   *Node
}

func NewList() *SingleList {
	return &SingleList{}
}
func (t *SingleList) GetItem(i int) (*Node, error) {
	//头节点为空，说明一个节点也没有
	if t.head == nil {
		return nil, errors.New("empty linked list")
	}
	if i < 0 || i > t.length-1 {
		return nil, errors.New("index is out of range")
	}
	//走完上面情况，这里就是只有一个节点的情况：头结点和尾节点都是它自己
	if t.tail == t.head {
		return t.head, nil
	}
	//剩下的情况至少有2个节点,开始遍历 三种情况：
	var seek *Node //要查找的节点
	//1.在头部；
	if i == 0 {
		seek = t.head
		return seek, nil
	}
	// 3.在尾部
	if i == t.length-1 {
		seek = t.tail
		fmt.Printf("Found node:%+v\n", seek)
		return seek, nil
	}
	current := t.head //当前节点指针
	var flag bool = false
	for j := 0; j < t.length; j++ {
		//找到了，2.在中间；
		if j == i {
			fmt.Printf("当前节点：%+v\n", current)
			seek = current
			flag = true
			break
		}
		current = current.next //继续往下寻找
	}
	if !flag { //取反才能进入判断if块里，表示flag==false，也就是没找到需要的节点 current==null
		return nil, errors.New("the node is not exist")
	} else {
		seek = current
	}
	return seek, nil
}

// 尾部弹出
func (t *SingleList) Popup() (*Node, error) {
	// 空链表 null
	if t.head == nil {
		return nil, errors.New("an empty linked list cannot be ejected")
	}
	//只有一个节点
	var tail = t.tail //接受断尾
	if t.head == t.tail {
		t.tail = nil
		t.head = nil
		t.length--
		return tail, nil
	}
	// 至少一个
	var node = t.head //当前节点
	for j := 0; j < t.length; j++ {
		if j < t.length-2 {
			//指针不断向下寻址
			node = node.next
		} else {
			tail = node.next
			//当指针偏移到倒数第三个节点(j==t.length-3)时，通过 node = node.text操作 node已经成了尾结点的前驱 previous
			t.tail = nil
			node.next = nil
			t.tail = node //新尾巴指向node
			break
		}
	}
	t.length--
	return tail, nil
}
func (l *SingleList) Append(item int) *SingleList {
	node := new(Node)
	node.item = item
	// node.next = nil 默认
	if l.head == nil {
		// 没有元素，头尾都是nil，但凡有一个元素，头尾都不是nil
		l.head = node
		l.tail = node
	} else {
		// 哪怕是有一个元素，append都是动尾巴
		l.tail.next = node // 手拉手1，修改当前尾巴的下一个
		l.tail = node      // 修改尾巴
	}
	l.length++ // 新增成功长度加1
	return l
}

func (t *SingleList) IterNodes() {
	if t.head == nil {
		panic("空空如也~")
	}
	//链表遍历靠指针寻址，不断下一跳
	p := t.head // 类型为 *Node节点指针
	for ; p != nil; p = p.next {
		fmt.Printf("%+v\n", p)
	}
}
