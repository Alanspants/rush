package link

import (
	"errors"
	"fmt"
	"strconv"
)

type Node struct { // 元素或结点
	item int
	next *Node // nil 后继
	prev *Node // nil 前驱
}

func (n *Node) String() string {
	//"1 --> 2"
	var s1, s2 string
	if n.next == nil {
		s1 = "null"
	} else {
		s1 = strconv.Itoa(n.next.item)
	}
	if n.prev == nil {
		s2 = "null"
	} else {
		s2 = strconv.Itoa(n.prev.item)
	}
	return fmt.Sprintf("%s <== %d ==> %s", s2, n.item, s1)
}

type LinkList struct {
	len  int
	head *Node // nil
	tail *Node // nil
}

func New() *LinkList {
	return &LinkList{}
}

func (l *LinkList) Len() int {
	return l.len // getter len字段只读
}

// 链表 Append Iternodes
func (l *LinkList) Append(value int) *LinkList {
	// 1 创建元素，next悬空nil
	fmt.Println("append called")
	node := new(Node)
	node.item = value
	// node.next = nil
	// node.prev = nil
	// 2 老尾巴指向node
	if l.tail == nil {
		// 2.1 一个元素都没有 len==0,l.head==nil,l.tail==nil
		l.head = node
	} else {
		// 2.2 至少一个元素
		l.tail.next = node
		node.prev = l.tail
	}
	// 3 更新尾巴，新尾巴
	l.tail = node
	l.len++
	return l
}

// 双向链表 Pop， Insert(index), Remove(index)
func (l *LinkList) Pop() (int, error) {
	// 相对于Append，把尾巴弹出,关注tail
	if l.tail == nil {
		// 条件l.len == 0, l.head == nil
		return 0, errors.New("Empty")
	}
	value := l.tail.item
	if l.head == l.tail {
		// 只有一个元素, l.len == 1, l.head == l.tail
		l.head = nil
		l.tail = nil
	} else {
		// 不止一个元素，该尾巴
		tail := l.tail
		prev := tail.prev
		prev.next = nil
		l.tail = prev
	}
	l.len--
	return value, nil
}

func (l *LinkList) Insert(index, value int) error {
	// 插入有2个版本，1、按index 2、按value值
	if index < 0 {
		return errors.New("not negative")
	}
	if index >= l.len {
		l.Append(value)
		return nil
	}
	// index >= 0 && index < l.len [0, l.len-1]
	// Go最好的实现，再写一遍迭代代码
	var current *Node // nil
	// var flag bool     // false
	for i, v := range l.Iternodes(false) {
		// fmt.Println(i, v, "###")
		if i == index {
			current = v
			// 下面写一大段代码，不好看，我们要把代码扁平化
			// flag = true
			break
		}
	}
	// if !flag {
	// 	// 没找到，说明不能在目前某个元素前插入，则使用尾巴追加
	// 	l.Append(value)
	// 	return nil
	// }
	// 找到了插入点，一定不是尾部追加，至少有1个元素, 从这行开始不动尾巴
	node := new(Node) // *Node
	node.item = value

	// 分2种：1、开头插入
	if current.prev == nil {
		// index == 0, current.prev == nil, l.head == current
		l.head = node
	} else {
		// 2 中间插入
		prev := current.prev // old
		prev.next = node
		node.prev = prev
	}
	current.prev = node
	node.next = current

	l.len++
	return nil
}

func (l *LinkList) Remove(index int) error {
	if l.tail == nil {
		// 条件l.len == 0, l.head == nil
		return errors.New("Empty")
	}
	if index < 0 {
		return errors.New("not negative")
	}
	if index >= l.len {
		return errors.New("out of index range")
	}
	// 至少有一个
	var current *Node // nil
	p := l.head
	for i := 0; i < l.len; i++ {
		if i == index {
			current = p
			break
		}
		p = p.next
	}
	// 一定找到了，开始考虑移除的情况
	prev := current.prev
	next := current.next
	if l.head == l.tail {
		// only 1 即是开头又是结尾
		// l.head == current && l.tail == current
		// current.prev==nil && current.next == nil
		// l.len == 1
		l.head = nil
		l.tail = nil
	} else if l.head == current {
		// prev == nil
		// 在开头，仅仅是开头，但不是结尾，多于一个元素
		l.head = next
		next.prev = nil
	} else if l.tail == current {
		// next == nil
		// 在尾部，仅仅是结尾，不是开头，多于一个元素
		prev.next = nil
		l.tail = prev
	} else {
		// 在中间，至少3个元素
		prev.next = next
		next.prev = prev
	}

	l.len--
	return nil
}

// func (l *LinkList) Extend(values ...int) {}

func (l *LinkList) Iternodes(reversed bool) []*Node {
	// 向前遍历？
	var p *Node
	r := make([]*Node, 0, l.len)
	if reversed {
		p = l.tail
	} else {
		p = l.head
	}
	for p != nil {
		// fmt.Println(p)
		r = append(r, p)
		if reversed {
			p = p.prev
		} else {
			p = p.next
		}
	}
	return r
}
