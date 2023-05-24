package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 单链表
type node struct {
	data int
	next *node
}

// type List *Node     //这么写method会报错 invalid receiver type List (pointer or interface type)

type linkList struct {
	head *node
}

func (l *linkList) append(n int) {
	newnode := &node{data: n, next: nil} //根据n的值构建新的node
	if l.head == nil {
		l.head = newnode //头指针为空，则将newnode当作首元节点挂到头指针下    //没有头节点，直接是  |头指针|->|首元节点|
	} else {
		p := l.head         //p指向首元节点
		for p.next != nil { //p不断后移，直至其指针域为空
			p = p.next
		}
		p.next = newnode //在指针域为空的p节点上挂上newnode
	}
}

func (l *linkList) iternodes() {
	if l.head == nil {
		fmt.Println("空链表")
	} else {
		p := l.head    //指向首元节点
		for p != nil { //判断节点是否为空
			fmt.Printf("-> %v ", p.data) //非空则打印并后移指针
			p = p.next
		}
	}
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var l linkList
	fmt.Print("初始链表为： ")
	l.iternodes() //初始为空链表
	for i := 0; i < 5; i++ {
		t := r.Intn(100)
		(&l).append(t)
		fmt.Println("对链表l尾部插入： ", t)
	}
	fmt.Print("插入后链表为：")
	(&l).iternodes()
}
