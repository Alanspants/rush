package main

import (
	"fmt"
	"time"
)

/*
1、8年前的今天上午9点30分
要求：
1. 毫秒时间戳是多少？
2. 格式化输出时间为 2005/09/10 21:35:40 +0800 的形式
3. 请问那天是周几？到那天，本年已经过了多少周？
4. 距离今天已经过了多少天了？
*/
func main() {
	//当前时间
	s := time.Now().Format("2006/01/02 15:04:05")
	//8年前的今天上午9点30分
	s1 := "2015/03/03 09:30:00"
	//时间解析，本地时间
	t1, err := time.Parse("2006/01/02 15:04:05", s)
	if err != nil {
		fmt.Println(err)
	}
	t2, err := time.Parse("2006/01/02 15:04:05", s1)
	if err != nil {
		fmt.Println(err)
	}
	//1. 毫秒时间戳是多少？
	fmt.Println(t2.UnixMilli())
	//	2. 格式化输出时间为 2005/09/10 21:35:40 +0800 的形式
	fmt.Println(t2.Local())
	//3. 请问那天是周几？到那天，本年已经过了多少周？
	_, week := t2.ISOWeek()
	fmt.Printf("那天是周%d,到那天，本年已经过了%d周\n", int(t2.Weekday()), week)

	//4. 距离今天已经过了多少天了？
	d := t1.Sub(t2)
	fmt.Printf("距离今天已经过了%d天了", int(d.Hours())/24)

}
