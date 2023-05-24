package link

import (
	"errors"
	"fmt"
	"strconv"
)

type OneWayNode struct {
	item int         // 节点存储的值
	next *OneWayNode // 指向下一个节点的地址，指针0值是nil
}

func (own *OneWayNode) String() string {
	var s string
	if own.next == nil {
		// 如果是尾部节点，own.next为nil，所以下一个指向为空，或者称own.next为悬空
		s = "null"
	} else {
		// 如果不是尾部节点，指向下一个的值为下一个节点的item
		// 因为item是int类型，所以由int转换为ASCII码
		s = strconv.Itoa(own.next.item)
	}

	// 拼接返回值，输出4个占位符，默认右对齐
	return fmt.Sprintf("|%4d -> %4s|", own.item, s)
}

type OneWayLinkedList struct {
	len  int         // 链表长度
	head *OneWayNode // 指针0值是nil
	tail *OneWayNode // 指针0值是nil
}

func NewOneWayLinkedList() *OneWayLinkedList {
	return &OneWayLinkedList{}
}

// Len 相当于getter方法，可以获取单向链表长度，不暴露len属性
func (owl *OneWayLinkedList) Len() int {
	return owl.len
}

func (owl *OneWayLinkedList) Append(value int) *OneWayLinkedList {
	// new()方法返回类型的指针
	own := new(OneWayNode)
	// 参数value，赋值给node的item属性，也就是节点存储的值
	own.item = value
	// 一个元素都没有 判断条件 len==0，owl.head==nil，owl.tail==nil都可以
	if owl.tail == nil {
		// 一个元素都没有，链表的head和tail都指向同一个元素
		// OneWayNode的next指向nil，因为指针的0值为nil，所以不用写
		owl.head = own
	} else {
		// 至少有一个元素，链表原来尾部OneWayNode的next指向新的own
		owl.tail.next = own
	}
	// 链表指向尾部的指针指向新的own if else公共语句向下提取
	owl.tail = own

	owl.len++
	// 返回追加后的链表自身，实现链式编程
	return owl
}

func (owl *OneWayLinkedList) IterOneWayNodes() {
	// 从头部的第一个节点赋值开始遍历
	own := owl.head
	// 如果节点不等于nil，说明有元素，就打印一下
	for own != nil {
		//fmt.Println(oneWayNode.item) // 遍历只打印节点存储的值
		// 为了便于理解，自定义打印输出的字符串
		fmt.Println(own)
		// 节点的下一个指向，赋给当前节点
		own = own.next
	}
}

type TwoWayNode struct {
	item int
	prev *TwoWayNode
	next *TwoWayNode
}

func (twn *TwoWayNode) String() string {
	var s1, s2 string
	if twn.prev == nil {
		s1 = "null"
	} else {
		s1 = strconv.Itoa(twn.prev.item)
	}

	if twn.next == nil {
		// 如果是尾部节点，own.next为nil，所以下一个指向为空，或者称own.next为悬空
		s2 = "null"
	} else {
		// 如果不是尾部节点，指向下一个的值为下一个节点的item
		// 因为item是int类型，所以由int转换为ASCII码
		s2 = strconv.Itoa(twn.next.item)
	}

	// 拼接返回值，输出4个占位符，默认右对齐
	return fmt.Sprintf("|%s<-%d->%s|", s1, twn.item, s2)
}

type TwoWayLinkedList struct {
	len  int
	head *TwoWayNode
	tail *TwoWayNode
}

func NewTwoWayLinkedList() *TwoWayLinkedList {
	return &TwoWayLinkedList{}
}

func (twl *TwoWayLinkedList) Len() int {
	return twl.len
}

func (twl *TwoWayLinkedList) Append(value int) *TwoWayLinkedList {
	twn := new(TwoWayNode)
	twn.item = value
	if twl.tail == nil {
		twl.head = twn
		twl.tail = twn
	} else {
		twl.tail.next = twn
		twn.prev = twl.tail
	}

	twl.tail = twn
	twl.len++

	return twl
}

func (twl *TwoWayLinkedList) IterTwoWayNodes(reversed bool) []*TwoWayNode {
	twns := make([]*TwoWayNode, 0, twl.len)
	var twn *TwoWayNode
	if reversed {
		twn = twl.tail
	} else {
		twn = twl.head
	}
	for twn != nil {
		twns = append(twns, twn)
		if reversed {
			twn = twn.prev
		} else {
			twn = twn.next
		}
	}

	return twns
}

func (twl *TwoWayLinkedList) Pop() (int, error) {
	if twl.head == nil {
		return 0, errors.New("two way linked list is empty")
	}
	twn := twl.tail
	if twl.head == twl.tail {
		twl.head = nil
		twl.tail = nil
	} else {
		twn.prev.next = nil
		twl.tail = twn.prev
	}

	twl.len--
	return twn.item, nil
}

func (twl *TwoWayLinkedList) Insert(index, value int) error {

	if index < 0 {
		return errors.New("index can not be negative")
	}

	if index >= twl.Len() {
		twl.Append(value)
		return nil
	}

	var current *TwoWayNode
	p := twl.head
	//for i, node := range twl.IterTwoWayNodes() {
	//	if index == i {
	//		current = node
	//		break
	//	}
	//}

	for i := 0; i < twl.len; i++ {
		if index == i {
			current = p
			break
		}
		p = p.next
	}

	twn := new(TwoWayNode)
	twn.item = value
	if current.prev == nil {
		twl.head = twn
	} else {
		twn.prev = current.prev
		twn.next = current
		current.prev.next = twn
	}

	twn.next = current
	current.prev = twn

	twl.len++
	return nil
}

func (twl *TwoWayLinkedList) Remove(index int) error {
	// 判断空
	if twl.head == nil && twl.tail == nil {
		return errors.New("two way linked list is empty")
	}
	// 判断索引小于0
	if index < 0 {
		return errors.New("index can not be negative")
	}
	// 判断索引越界
	if index >= twl.len {
		return errors.New("index out of range")
	}

	// 走到这，链表至少有一个元素
	var current *TwoWayNode
	p := twl.head
	for i := 0; i < twl.len; i++ {
		if index == i {
			current = p
			break
		}
		p = p.next
	}

	if current.prev == nil && current.next == nil {
		// 只有一个元素
		twl.head = nil
		twl.tail = nil
	} else if current.prev == nil {
		// 多余一个元素，元素在开头
		current.next.prev = nil
		twl.head = current.next
	} else if current.next == nil {
		// 多余一个元素，元素在开头
		current.prev.next = nil
		twl.tail = current.prev
	} else {
		// 多余两个元素，元素在中间
		current.prev.next = current.next
		current.next.prev = current.prev
	}

	twl.len--
	return nil
}
