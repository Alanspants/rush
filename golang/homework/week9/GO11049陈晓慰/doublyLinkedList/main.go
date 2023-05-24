package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	data int
	prev *node
	next *node
}

type linkList struct {
	length int
	head   *node
	tail   *node
}

func (l *linkList) append(n int) {
	newNode := &node{data: n}
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
		l.length += 1
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
		l.length += 1
	}
}

func (l *linkList) iternodes() {
	p := l.head
	if p == nil {
		fmt.Println("empty linkList")
	} else {
		for p != nil {
			fmt.Printf("-> %v ", p.data)
			p = p.next
		}
	}

}

func (l *linkList) pop() (int, error) {
	if l.length == 0 { //或者判断l.length是否为0
		return 0, errors.New("can't pop in a empty linkList")
	} else if l.length == 1 {
		n := l.head.data
		l.head = nil
		l.tail = nil
		l.length--
		return n, nil
	} else {
		p := l.tail
		l.tail = l.tail.prev
		l.tail.next = nil
		p.prev = nil
		l.length--
		n := p.data
		return n, nil
	}
}

func (l *linkList) insert(data, index int) error {
	newnode := &node{data: data}
	switch {
	case l.length < index || index < 0:
		return errors.New("index out of range")
	case l.length == 0 && index == 0:
		l.head = newnode
		l.tail = newnode
		l.length++
	case index == 0:
		newnode.next = l.head
		l.head.prev = newnode
		l.head = newnode
		l.length++
	case l.length == index:
		l.tail.next = newnode
		newnode.prev = l.tail
		l.tail = newnode
		l.length++
	case index < l.length:
		p := l.head
		for i := 1; i < index; i++ {
			p = p.next
		}
		p.next.prev = newnode
		newnode.next = p.next
		p.next = newnode
		newnode.prev = p
		l.length++
	}
	return nil
}

// 注意receiver必须为指针，否则虽然改了链表里的内容，但因为是结构体复制，遍历的时候还是会用指向原来head的linklist去遍历，index=0异常
func (l *linkList) remove(index int) error {
	switch {
	case l.length == 0:
		return errors.New("empty linkList")

	case index < 0 || index >= l.length:
		return errors.New("index out of range")
	case l.length == 1:
		l.head = nil
		l.tail = nil
		l.length--
	case index == 0:
		l.head = l.head.next
		l.head.prev = nil
		l.length--

	case index == l.length-1:
		l.tail = l.tail.prev
		l.tail.next = nil
		l.length--
	default:
		p := l.head
		for i := 0; i < index-1; i++ {
			p = p.next
		}
		p.next = p.next.next
		p.next.prev = p
		l.length--
	}
	return nil
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
	fmt.Println()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~")

	// for j := l.length; j >= 0; j-- {
	// 	i, err := l.pop()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(i, l.length)
	// 	}
	// }
	err := l.insert(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		l.iternodes()
	}

	err = l.remove(3)
	if err != nil {
		fmt.Println(err)
	} else {
		l.iternodes()
	}
}
