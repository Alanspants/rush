package main

import (
	"fmt"
)

//链表是个容器 内部装这一个个节点对象
type LinkList struct {
	LinkSlice []*Lnode
}

//将链表节点追加到切片中
func (l *LinkList) append(newNode *Lnode) {

	l.LinkSlice = append(l.LinkSlice, newNode)
}

//遍历切片中的元素
func (l *LinkList) iternodes() {
	for i, v := range l.LinkSlice {
		fmt.Printf("第%v个切片,元素:%#v\n", i, v)
	}

}

type Lnode struct {
	data int
	next *Lnode
}

//构造函数 创建链表节点
func NewLnode(data int) (dlNode *Lnode) {
	return &Lnode{data: data}
}

//尾插法 插入链表节点
func TailInsert(head *Lnode, newNode *Lnode) {
	temp := head
	for temp.next != nil {
		temp = temp.next //指针后移 知道指向最后一个链表节点
	}
	temp.next = newNode
}

//单链表的测试
func testLinkList() {
	//创建切片用于存放*Lnode
	sliceLnode := make([]*Lnode, 0, 10)
	//创建结对象实例 用于存放[]*Lnode
	linklist1 := LinkList{sliceLnode}

	//创建链表头结点
	head := NewLnode(1)
	//将头结点放入切片中
	linklist1.append(head)

	//创建节点1
	lnode1 := NewLnode(2)
	//将链表节点尾插法进行连接
	TailInsert(head, lnode1)
	//将节点append到linklist1上
	linklist1.append(lnode1)

	//创建节点2
	lnode2 := NewLnode(3)
	//将链表节点尾插法进行连接
	TailInsert(head, lnode2)
	//将节点append到linklist1上
	linklist1.append(lnode2)

	//遍历链表
	linklist1.iternodes()
}

//双向链表
type DLinkList struct {
	DLinkSlice []*DLNode //切片数组用于存放 *DLNode类型
}

//将链表节点 追加到 结构体的切片上
func (l *DLinkList) append(newDLNode *DLNode) {
	l.DLinkSlice = append(l.DLinkSlice, newDLNode)
}

//对双向链表结构体对象的遍历
func (l *DLinkList) iternodes() {
	for i, v := range l.DLinkSlice {
		fmt.Printf("切片中第%v个,元素:%#v\n", i, v)
	}
}

//从切片中删除最后一个元素的数据 并返回这个数据
func (l *DLinkList) pop() (res int) {

	res = pop(l.DLinkSlice[0])

	//切片的删除
	l.DLinkSlice = append(l.DLinkSlice[:len(l.DLinkSlice)-1])

	return res
}

//按照索引插入元素  按照索引在切片中插入元素DLnode, index是要插入的位置 data是要插入的数据
func (l *DLinkList) insert(index int, data int) {
	//创建以个新节点
	newDLNode := NewDLNode(data)
	//执行链表内的插入
	insert(l.DLinkSlice[0], newDLNode, index)
	//DLinkSlice []*DLNode
	//更新[]*DLnode中的元素
	l.DLinkSlice = append(l.DLinkSlice[:index], append([]*DLNode{newDLNode}, l.DLinkSlice[index:]...)...)

}

type DLNode struct {
	data int     //数据域
	per  *DLNode //前指针
	next *DLNode //后指针
}

//链表尾插
func TailInsertDLNode(head *DLNode, newDLNode *DLNode) {
	//辅助节点 用于找到表尾指针
	temp := head

	for temp.next != nil {
		temp = temp.next
	}
	//找到了表尾指针
	fmt.Println("最后一个元素", temp)
	temp.next = newDLNode
	newDLNode.per = temp

}

//构造函数 用于生成 双向链表的节点
func NewDLNode(data int) *DLNode {
	return &DLNode{data: data}
}

//弹出链表中最后一个元素 删除并返回器数据
func pop(head *DLNode) (out int) {

	temp := head
	prior := head.per

	if head == nil {
		fmt.Println("链表为空不可弹处-99")
		return -99
	}
	for temp.next != nil {
		temp = temp.next
		prior = temp.per
	}
	//找到了链表表尾
	//截断链表最后一个元素
	prior.next = nil
	temp.per = nil

	//将截断的输出
	return temp.data
}

//按照索引插入元素
func insert(head *DLNode, newDLNode *DLNode, index int) {

	length := 0
	//遍历双向链表找到插入位置

	temp := head
	if temp != nil {
		length++
	}
	prior := head.per
	for temp.next != nil {

		temp = temp.next
		length++
		prior = temp.per
		if index == length {
			//找到了插入位置
			newDLNode.next = temp.next
			temp.next.per = newDLNode
			newDLNode.per = prior
			prior.next = newDLNode

		}
	}

}

//用于测试双向链表的功能
func testDLinkList() {

	//make一个切片用于存放*DLNode
	dLinkSlice := make([]*DLNode, 0, 10)
	//创建一个双向链表对象
	dllist1 := DLinkList{dLinkSlice}

	//创建双链表的头结点
	head := NewDLNode(1)
	//将双想链表头结点 放入切片中
	dllist1.append(head)

	//创建双链表结点2
	dlNode1 := NewDLNode(2)
	//将双想链表头结点 放入切片中
	TailInsertDLNode(head, dlNode1)
	dllist1.append(dlNode1)

	//创建双链表结点2
	dlNode2 := NewDLNode(3)
	//将双想链表头结点 放入切片中
	TailInsertDLNode(head, dlNode2)
	dllist1.append(dlNode2)

	//遍历对象
	//dllist1.iternodes()

	fmt.Println("pop之前的对象")
	dllist1.iternodes()

	res := dllist1.pop()
	//pop弹处链表最后一个元素
	fmt.Println("尾部弹出元素后的链表对象", "弹出的值是=", res)
	dllist1.iternodes()

	fmt.Println("在index索引出插入一个新节点前")
	dllist1.iternodes()
	//在index索引出插入一个新节点
	dllist1.insert(2, 6)
	fmt.Println("在index索引出插入一个新节点后")
	dllist1.iternodes()

}
func main() {
	//单链表的测试
	testLinkList()

	//双链表的测试
	//testDLinkList()
}
