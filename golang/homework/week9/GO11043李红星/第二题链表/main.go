package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Node struct {
	item int
	next *Node //nil
	prev *Node //nil
}

type DoubleLinkList struct {
	len  int
	head *Node
	tail *Node
}

func NewDoubleLinkList() *DoubleLinkList {
	return &DoubleLinkList{}
}

func (n *Node) String() string {
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
	return fmt.Sprintf("%s<--%d--->%s", s2, n.item, s1)
}

//遍历链表取出各个元素
func (l *DoubleLinkList) Iternodes() []*Node {
	//向后遍历 找链表头节点
	p := l.head
	//复用这个for循环给insert的方法用
	nodelist := make([]*Node, 0, l.len)
	for p != nil {
		nodelist = append(nodelist, p)
		//insert方法的时候暂时注释掉该打印便于查看
		//fmt.Println(p)
		p = p.next
	}
	return nodelist
	//向后遍历 找链表尾节点
	//p = l.tail
	//for p != nil {
	//	fmt.Println(p)
	//	p = p.prev
	//}
}

//追加元素
func (l *DoubleLinkList) Append(value int) *DoubleLinkList {
	//1、构造node
	node := new(Node)
	node.item = value
	//node.next=nil  初始化默认为nil
	//node.prev=nil  初始化默认为nil

	//2、老尾巴指向node
	if l.tail == nil {
		//2.1一个元素也没有，头尾都是当前node
		l.head = node
		//l.tail = node //至少有一个的时候也会挪尾巴，抽出来
		//node.prev = nil  //l.tail,prev ==nil  //默认为nil 一个没有追加了一个，此时prev next都还是nil
		//node.next = nil  //l.tail.next ==nil  //默认不写为nil
	} else {
		//至少有一个的时候，head头不动，前一个的next指下一个，下一个的prev指前一个尾巴、其他不动
		//l.head  //尾部追加的时候头不动
		//l.tail.prev //不动
		l.tail.next = node
		node.prev = l.tail
		//node.next = nil   //默认为空
		//l.tail = node  //在没元素的时候追加一个也会挪尾巴，抽出来
	}
	//更新尾巴
	l.tail = node
	l.len++
	return l
}
func (l *DoubleLinkList) Pop() error {
	//1、没有任何元素的时候弹出错误，没有元素 即l.len==0 或 l.tail == nil 或 l.head==nil
	if l.len == 0 {
		return errors.New("empty")
	} else if l.head == l.tail {
		//只有一个元素，l.head==0xc000004090 l.tail==0xc000004090,所以上面在没有元素的时候 不推荐l.head==l.tail判断
		//上面已经把0个元素过滤掉了，所以这里l.head==l.tail做判断不会出现0个元素的情况
		//移除节点后 没有了节点 所以head和tail都是nil
		l.head = nil
		l.tail = nil
	} else {
		//不止一个元素的时候弹出
		//1、断开手拉手，即前一个(l.tail.prev)的next = nil
		//2、更新尾巴 即 l.tail = 前一个节点
		//其他不必要的手拉手 不管 即 l.tail.prev =nil l.tail.next=nil
		//前面的节点不动，更新len
		l.tail.prev.next = nil
		l.tail = l.tail.prev

	}
	l.len--
	return nil
}

func (l *DoubleLinkList) Insert(index, value int) error {
	//如果索引为负数
	if index < 0 {
		return errors.New("invalid position")
	}
	//index>0
	//定义当前节点
	var current *Node
	//Iternodes至少有一个元素的时候
	for i, node := range l.Iternodes() {
		//找到当前节点
		if index == i {
			current = node
			break
		}
		//Iternodes至少有一个元素的时候,如果遍历完Iternodes还没有和insret的索引匹配证明索引超界,那也追加
		//如果Iternodes 没有元素的时候 也会执行尾部追加 append会写两次,所以此处注释掉挪外面
		//else {
		//	l.Append(value)
		//}
	}
	//如果没找到,可能超界;也可能这个时候for循环也没走即Iternodes里没有任何元素,就尾部追加
	if current == nil {
		l.Append(value)
		return nil
	}
	//找到了插入点,一定不是尾部追加
	//分两种 1是开头插入 2是中间插入 都不会动他的尾巴
	node := new(Node)
	node.item = value
	//开头插入的条件 index=0 或current.prev = nil 或 l.head = nil
	if index == 0 {
		//当前节点的prev指到新增的节点
		current.prev = node
		//头 改成新增节点
		l.head = node
		//新头的下一个指到当前节点
		l.head.next = current //node.next = current
	} else {
		node.prev = current.prev
		node.prev.next = node
		current.prev = node
		current.prev.next = current //node.next = current

	}
	l.len++
	fmt.Println(current)
	return nil
}

func main() {
	dll := NewDoubleLinkList()
	l := dll.Append(100).Append(200).Append(300)
	fmt.Println(l)
	l.Iternodes()
	fmt.Println("=====Pop======")
	p := l.Pop()
	fmt.Println(p)
	l.Iternodes()
	fmt.Println("------insert------")
	l.Insert(1, 50)
	fmt.Println(l.Iternodes())
}
