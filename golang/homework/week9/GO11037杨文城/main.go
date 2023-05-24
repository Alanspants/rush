package main

import (
	"fmt"
	"math"
	"time"
)

func work1() {
	// 1、8年前的今天上午9点30分
	// 要求：
	// 1. 毫秒时间戳是多少？
	// 2. 格式化输出时间为 2005/09/10 21:35:40 +0800 的形式
	// 3. 请问那天是周几？到那天，本年已经过了多少周？
	// 4. 距离今天已经过了多少天了

	// 当前时间
	t1 := time.Now()
	fmt.Println("今天日期", t1)
	// 获取月日
	tdate := t1.Format("01-02")
	// 获取8年前的年份
	y1 := t1.Year() - 8
	// 组装字符串
	y1Str := fmt.Sprintf("%d-%s 09:30:00", y1, tdate)
	// 字符串转日日期
	fmt.Println("8年前上午9点30分日期字符串", y1Str)

	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", y1Str, tz)
	fmt.Println("日期对象", t2)
	// 输出毫秒
	fmt.Println("毫秒", t2.UnixMilli())
	fmt.Println(t2.Format("2006/01/02 15:04:05 -0700"))
	fmt.Println(t2.Weekday())
	_, w1 := t2.ISOWeek()
	fmt.Printf("本年已经过了%d周 \n", w1)
	d1 := t1.Sub(t2)

	fmt.Println("距离的天数", math.Ceil(d1.Seconds()/86400))
}

func main() {
	work1()
	// 链表
	work2()
}

func work2() {
	fmt.Println("单链表开始")
	L1 := NewLinkedList(1)
	L1.append(2)
	L1.append(3)
	L1.iternodes()
	fmt.Println("双链表开始")
	dl := NewDoubleLinkedList(1)
	dl.append(2)
	dl.append(3)
	dl.iternodes()
	fmt.Println("弹出最后一个元素")
	dl.pop()
	dl.iternodes()
	fmt.Println("插入一个元素")
	dl.insert(0, 4)
	dl.iternodes()
	fmt.Println("删除一个元素")
	dl.remove(1)
	dl.iternodes()
}

type LinkedList struct {
	next *LinkedList
	data int
}

func NewLinkedList(data int) *LinkedList {
	ll := LinkedList{data: data}
	return &ll
}

func (ll *LinkedList) append(data int) {
	//头结点处理一下
	for {
		if ll.next != nil {
			// 找到最后一个元素
			ll = ll.next
		} else {
			ll.next = NewLinkedList(data)
			break
		}
	}
}

func (ll *LinkedList) iternodes() {
	for {
		fmt.Printf("元素地址 %p 数据 %d  next: %p \n", ll, ll.data, ll.next)
		if ll.next != nil {
			ll = ll.next
		} else {
			break
		}
	}
}

type DoubleLinkedList struct {
	next *DoubleLinkedList
	prev *DoubleLinkedList
	data int
}

func NewDoubleLinkedList(data int) *DoubleLinkedList {
	dl := DoubleLinkedList{data: data}
	return &dl
}

func (dl *DoubleLinkedList) append(data int) {

	for {
		if dl.next != nil {
			// 找到最后一个元素
			dl = dl.next
		} else {
			tmpDl := NewDoubleLinkedList(data)
			tmpDl.prev = dl
			dl.next = tmpDl
			break
		}
	}
}

func (dl *DoubleLinkedList) iternodes() {
	for {
		fmt.Printf("元素地址 %p %d  pref %p  next %p \n", dl, dl.data, dl.prev, dl.next)
		if dl.next != nil {
			dl = dl.next
		} else {
			break
		}
	}
}

func (dl *DoubleLinkedList) pop() {
	for {
		if dl.next != nil {
			dl = dl.next
		} else {
			dl.prev.next = nil
			break
		}
	}
}

func (dl *DoubleLinkedList) insert(idx int, data int) {
	index := 0
	for {
		if index == idx {
			inserDl := NewDoubleLinkedList(data)
			// 前一个修改
			if dl.prev != nil {
				dl.prev.next = inserDl
				inserDl.prev = dl.prev
				inserDl.next = dl
				dl.prev = inserDl

			} else {
				// 从头插入特殊处理
				inserDl.prev = dl
				inserDl.next = dl.next
				inserDl.data = dl.data
				if dl.next != nil {
					dl.next.prev = inserDl
				}
				dl.next = inserDl
				dl.data = data
			}
			break
		}

		if dl.next != nil {
			dl = dl.next
		} else {
			panic("索引超出范围")
		}
		index++
	}
}

func (dl *DoubleLinkedList) remove(idx int) {
	index := 0
	for {
		if index == idx {

			if dl.prev == nil {
				// 开头删除
				if dl.next != nil {
					dl.data = dl.next.data
					dl.next = dl.next.next
					dl.next.prev = dl
				} else {
					panic("至少保留一个元素")
				}

			} else {
				// 中间删除
				dl.prev.next = dl.next
				if dl.next != nil {
					dl.next.prev = dl.prev
				}
			}

			break
		}

		if dl.next != nil {
			dl = dl.next
		} else {
			panic("索引超出范围")
		}
		index++
	}
}
