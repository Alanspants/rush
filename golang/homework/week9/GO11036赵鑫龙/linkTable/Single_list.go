package linkTable

import (
	"fmt"
	"strconv"
)

type Node struct {
	value int
	next  *Node
	front *Node
}

type LinkList struct {
	len  int
	head *Node
	tail *Node
}

func New() *LinkList {
	return &LinkList{}
}
func (l *LinkList) Len() int {
	return l.len
}

func (n *Node) String() string {
	var next string
	if n.next == nil {
		next = "null"
	} else {
		next = strconv.Itoa(n.next.value)
	}

	return fmt.Sprintf("%d==>%s", n.value, next)
}

func (l *LinkList) Append(value int) *LinkList {
	node := new(Node)
	node.value = value
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	return l
}

func (l *LinkList) IterNodes() {
	p := l.head
	for p != nil {
		fmt.Println(p)
		p = p.next
	}
}
