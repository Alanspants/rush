package main

import "fmt"

//每一个结点是一个独立的对象，结点对象有内容，还知道下一跳是什么。
type Node struct {
	data interface{} //内容
	Next *Node       //下一跳
}

type LinkList struct {
	Len int
	Ptr []*Node
}

//遍历查看链表数据
func (ll *LinkList) iternodes() {

	for _, v := range ll.Ptr {
		fmt.Printf("iternodes v: %v\n", v)
	}
}

//尾部添加节点
func (ll *LinkList) append(n *Node) {
	// fmt.Println(&n)

	if ll.Len == 0 {
		ll.Ptr = append(ll.Ptr, n)
		ll.Len++
		return
	}
	for _, v := range ll.Ptr {
		if v.data == n.data {
			return
		}
	}
	ll.Ptr = append(ll.Ptr, n)
	ll.Len++
	return
}

func Link() {
	LinkList := new(LinkList)

	n1 := new(Node)
	n1.data = 111
	LinkList.append(n1)
	n2 := new(Node)
	n2.data = 222
	LinkList.append(n2)
	LinkList.append(n2)
	fmt.Printf("LinkList.Len: %v\n", LinkList.Len)
	LinkList.iternodes()

	// time()
}
