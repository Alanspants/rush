package main

import (
	"fmt"
	"time"
)

func main() {
	//设置时区
	l, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(l)
	}
	//计算八年前
	t := time.Now().AddDate(-8, 0, 0)
	// 8年前的今天上午9点30分
	t2 := time.Date(t.Year(), t.Month(), t.Day(), 9, 30, 0, 0, l)
	fmt.Println(t2)
	t3, _ := time.Parse("20060102 15:04:05 -0700", fmt.Sprintf("%s 09:30:00 %s", t.Format("20060102"), t.Format("-0700")))
	fmt.Println(t3)
	// 1. 毫秒时间戳是多少？
	fmt.Println(t2.UnixMilli())
	// 2. 格式化输出时间为 2005/09/10 21:35:40 +0800 的形式
	fmt.Printf("%+v\n", t2.Format("2006/01/02 15:04:05 -0700"))
	// 3. 请问那天是周几？到那天，本年已经过了多少周？
	fmt.Println(t2.Weekday())
	year, week := t2.ISOWeek()
	fmt.Println(year, week)
	// 4. 距离今天已经过了多少天了？
	d := time.Now().Sub(t2)
	fmt.Println(d, int(d.Hours()/24))
	d = time.Since(t3)
	fmt.Println(d, int(d.Hours()/24))
}
