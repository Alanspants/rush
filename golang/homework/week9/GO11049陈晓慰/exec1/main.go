package main

import (
	"fmt"
	"time"
)

// 因time.Weekday的String() string方法打印的是英文格式的周几，通过新建类型及方法重新输出为中文周几
type myWeekday time.Weekday

var chinaDayNames = []string{
	"周日",
	"周一",
	"周二",
	"周三",
	"周四",
	"周五",
	"周六",
}

func (d myWeekday) String() string {
	return chinaDayNames[d]
}

/*
1、8年前的今天上午9点30分
要求：

	(1) 毫秒时间戳是多少？
	(2) 格式化输出时间为2005/09/10 21:35:40 +0800的形式
	(3) 请问那天是周几？到那天，本年已经过了多少周？
	(4) 距离今天已经过了多少天了？
*/
func EightYearAgo() {
	//（1）先获取并打印当前时间now
	now := time.Now()
	fmt.Printf("今天是: %v\n", now.Format("2006-01-02"))
	//（2）构造8年前的今天上午9点30分实例targetTime
	targetTime := time.Date(now.Year()-8, now.Month(), now.Day(), 9, 30, 0, 0, time.Local)
	fmt.Printf("8年前的今天上午9点30分是: %v\n", targetTime)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("其毫秒时间戳是: %v\n", targetTime.UnixMilli())
	fmt.Printf("以2005/09/10 21:35:40 +0800的形式格式化输出为： %v\n", targetTime.Format("2006/01/02 15:04:05 -0700"))
	year, week := targetTime.ISOWeek()
	fmt.Printf("那天是%v，到那天，%v年已经过了%v周\n", myWeekday(targetTime.Weekday()), year, week)
	fmt.Printf("那天距离今天已经过了%v天\n", int(now.Sub(targetTime)/(24*time.Hour)))

}

type node struct {
	data int
	next *node
}

type linklist *node

type head struct {
	head *node
}

func createNode(data int) *node {
	return &node{data: data}
}

func main() {
	//定义链表
	var L linklist
	fmt.Printf("L: %+v\n", L)
	var p *node
	var n node
	fmt.Printf("p: %+v\n", p)
	fmt.Printf("n: %+v\n", n)
	// a := createNode(1)
	// b := createNode(2)
	// fmt.Println(a, b)

	// EightYearAgo()

}
