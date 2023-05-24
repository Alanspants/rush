package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Element struct {
	previous *Element
	next     *Element
	data     int
}

func (t *Element) String() string {
	var previous, next string
	if t.previous == nil {
		previous = "null"
	} else {
		previous = fmt.Sprintf("%d", t.previous.data)
	}
	if t.next == nil {
		next = "null"
	} else {
		next = fmt.Sprintf("%d", t.next.data)
	}
	return fmt.Sprintf("%s <== %d ==> %s", previous, t.data, next)
}

type LinkedList struct {
	head *Element
	tail *Element
	len  int
	err  error
}

func New() *LinkedList {
	return &LinkedList{}
}
func (t *LinkedList) append(data int) *LinkedList {
	fmt.Println("append called")
	// 创建新节点
	node := new(Element)
	node.data = data //赋值
	// 追加有两种情况，有无节点：2、3、可以合并成一种，但凡有一个元素都要动尾巴
	// 1、空链表 t.head==t.tail == nil || t.length==0
	// 2、有一个节点 t.head==t.tail !== nil || t.length==1
	// 3、至少一个元素 尾结点后继指向新节点:t.tail.next=node，新节点前驱指向尾结点:node.previous=t.tail，改尾巴为新节点:t.tail=node
	if t.tail == nil { //无
		t.tail = node
		t.head = node // =t.tail第一次运行会报空指针
	} else { //有
		t.tail.next = node
		node.previous = t.tail
		t.tail = node
	}
	t.len++
	t.err = nil
	return t
}
func (t *LinkedList) iternodes(reverse bool) []*Element {
	var cur *Element
	nodes := make([]*Element, 0, t.len)
	if reverse {
		cur = t.tail
	} else {
		cur = t.head
	}
	for cur != nil {
		nodes = append(nodes, cur)
		if reverse {
			// 指针前摇
			cur = cur.previous
		} else {
			// 指针后摆
			cur = cur.next
		}
		fmt.Printf("%+v %[1]p\n", cur)
	}
	return nodes
}
func (t *LinkedList) pop() (*Element, *LinkedList) {
	fmt.Println("popup·····························长度：", t.len)
	// 1.空链表
	if t.tail == nil {
		t.err = errors.New("the empty linked list cannot be ejected")
		return nil, t
	}
	// 2.只有一个元素 首尾相同
	tail := t.tail
	if t.head == t.tail {
		t.head = nil
		t.tail = nil
	} else { // 3.至少一个节点 尾结点的特点是：后继next是空的，新尾巴（旧尾巴的前驱）也要有这个特点，即：t.tail.previous.next=nil; t.tail=t.tail.previous;
		t.tail = t.tail.previous //当前尾结点取前驱
		t.tail.next = nil        //尾结点（.tail.previous）后继置空，才满足尾结点特性
	}
	t.len--
	return tail, t
}
func (t *LinkedList) InsertByIndex(index, value int) *LinkedList {
	if index < 0 {
		t.err = errors.New("the index cannot be negative")
		return t
	}
	//创建新节点
	element := new(Element)
	// 赋值
	element.data = value
	// 1、index大于等于length则追加 append;
	// 2、索引==0;
	// 3、小于长度t.len就插入,范围前开后闭(0, t.len-1] 重新手拉手
	if index > t.len-1 {
		t.append(value)
	} else if index == 0 { //在头前插入,换头
		t.head.previous = element
		element.next = t.head
		t.head = element
	} else {
		p := t.head
		for i := 1; i < t.len; i++ {
			if i == index {
				element.next = p //element后继指向p
				break
			}
			p = p.next //游标卡尺放这里是为了避免放if判断上面index==0的情况发生即插入head前，p直接向下移成了p.next不再是头结点，这样element就插错了位置。
		}
		element.previous = p.previous //新节点element的前驱设置为原来p的前驱。放在循环体外防止指针不断往下偏移出现 p.previous空指针
		// ①新节点先占位
		p.previous.next = element //p前驱的后继位置给element。这行可以和上面对调，但绝不可以和下面对调！
		// ②当前节点p才能靠后往前伸手拽e
		p.previous = element //这行代码只能放在p.previous.next = element后面执行，否则：p.previous空指针！！
	}
	t.len++
	return t
}
func (t *LinkedList) InsertByValue(value int) *LinkedList {
	ele := new(Element)
	ele.data = value
	//空链表直接追加
	if t.tail == nil {
		//这里必须用return,不能用 t = t.append(value)，否则，t接受了append的返回值就继续插入导致最后节点多加一个
		return t.append(value)
	}
	var node *Element
	curP := t.head
	for curP != nil {
		if curP.data == value {
			node = curP
			break
		}
		curP = curP.next
	}
	//没找到value就追加
	if node == nil {
		return t.append(value) //这里必须用return,不能为了扁平化统一返回而用 t = t.append(value)，否则，t接受了append的返回值就继续插入导致最后节点多加一个
	}
	// 找到了 重新拉手
	if t.head == t.tail { //只有一个节点 首尾相同
		ele.next = t.head
		t.head.previous = ele
		t.head = ele
		t.len++
		return t
	}
	ele.previous = curP.previous
	ele.next = curP
	curP.previous.next = ele
	curP.previous = ele
	t.len++
	return t
}
func (t *LinkedList) remove(index int) *LinkedList {
	// 空链表
	if t.head == nil {
		t.err = errors.New("an empty linked list cannot remove elements")
		return t
	}
	if index < 0 {
		t.err = errors.New("the index cannot be negative")
		return t
	}
	if index > t.len-1 {
		t.err = errors.New("the index is out of range")
		return t
	}
	//只有一个节点
	if t.head == t.tail {
		if index == 0 {
			t.head = nil
			t.tail = nil
			t.len--
		} else {
			t.err = errors.New("index not found")
		}
		return t
	}
	//至少一个
	p := &Element{}
	current := t.head
	for i := 0; i < t.len; i++ {
		if i == index {
			p = current
			break
		}
		current = current.next
	}
	if current.previous == nil { //在开头
		t.head = p.next
		p.next.previous = nil //头结点前驱为空
	} else if current.next == nil { //在结尾
		t.tail = p.previous
		p.previous.next = nil //尾结点后继为空
	} else { //在中间
		// 当前节点前驱的后继指向当前节点的后继
		p.previous.next = p.next
		// 当前节点后继的前驱指向当前节点的前驱
		p.next.previous = p.previous
	}
	t.len--
	return t
}
func (t *LinkedList) GetNode(index int) *Element {
	if t.tail == nil {
		t.err = errors.New("empty linked list")
		return nil
	}
	if index < 0 || index > t.len-1 {
		t.err = errors.New("index is out of range")
		return nil
	}
	var current *Element
	p := t.head
	for i := 0; i < t.len; i++ {
		if i == index {
			current = p
		}
		p = p.next
	}
	return current
}
func (t *LinkedList) ExtendRemove(ints ...int) error {
	if t.head == nil {
		return errors.New("nodes cannot be deleted from an empty linked list")
	}
	m := make(map[int]any, len(ints))
	for _, v := range ints {
		//去重
		m[v] = nil
	}
	var str string
	keys := make([]int, 0, len(m))
	for k := range m {
		if k < 0 || k >= t.len { //超出范围的索引丢弃
			str += string(rune(k)) + ","
			fmt.Printf("%s\n", str)
			delete(m, k) //拼完后删除k，不然又进入了下面的循环
		} else {
			keys = append(keys, k)
		}
	}
	sort.Ints(keys) //排序
	p := t.head
	for k := range keys {
		if k == 0 { //在头部
			if t.head == t.tail { //只有一个节点 删完拉倒
				t.head = nil
				t.tail = nil
				return nil //所以，直接返回
			}
			//索引在头部，但有多个节点
			t.head = p.next       //当前节点在头部 换头，头结点变成当前节点的下一跳p.next
			p.next.previous = nil //头结点(p.next)的前驱是空
		} else if k == t.len-1 { //在尾部
			t.tail.previous.next = nil
			t.tail = nil
		} else { //在中间
			if p.previous != nil {
				p.previous.next = p.next
			}
			p.next.previous = p.previous
		}
		p = p.next
	}
	if str != "" {
		str = "Index: " + strings.TrimSuffix(str, ",") + " are out of list scope"
		return errors.New(str)
	} else {
		return nil
	}
}
