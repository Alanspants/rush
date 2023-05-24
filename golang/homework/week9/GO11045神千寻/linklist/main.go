package main

import "fmt"

func main() {
	// 构造一个单链表
	ll := NewList()
	ll.Append(1).Append(2).Append(3)
	fmt.Printf("%#v\n", ll) //只有1个节点时,头和尾指针相同:&main.SingleList{length:1, head:(*main.Node)(0xc000058270), tail:(*main.Node)(0xc000058270)}
	fmt.Println("++++++++++++++++++++++++++++++++++++单链表开始遍历+++++++++++++++++++++++++++++++++++++")
	ll.IterNodes()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~单链表出栈~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("%+v\n", ll)
	fmt.Println(ll.Popup())
	fmt.Println(ll.Popup())
	fmt.Println(ll.Popup())
	fmt.Printf("%+v\n", ll)
	fmt.Println("====================================================双链表开始遍历======================================================")
	// 开始操作双向链表
	sl := New()
	sl.append(10).append(11).append(12)
	fmt.Printf("%+v\n", sl.iternodes(true))
	fmt.Println("长度：", sl.len)
	// sl.InsertByIndex(2, 8).InsertByIndex(3, 7).InsertByIndex(4, 6)
	fmt.Println("—————————————————————————————————————————————————————按值插入双向链表—————————————————————————————————————————————————————")
	sl.InsertByValue(0).InsertByValue(1).InsertByValue(2)
	fmt.Printf("%v\n", sl.iternodes(false))
	fmt.Println("长度：", sl.len)
	// 开始出栈
	/* sl.pop()
	e, r := sl.pop()
	fmt.Printf("%+v\n", e)
	fmt.Printf("%+v err:%s\n", r, r.err) */
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Printf("%+v\n", sl.ExtendRemove(-1, 0, 1, 10, 2, 2))
	fmt.Printf("%v\n", sl.iternodes(false))
	fmt.Println("长度：", sl.len)
}
