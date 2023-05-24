package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	//format := now.Format("2006/01/02 15:04:05.000 -0700")
	//year := now.Year() - 8
	//monthAndDay := format[4:10]
	//timeStr := fmt.Sprintf("%d%s %s", year, monthAndDay, "09:30:00") //八年前的今天上午9点30分

	date := now.AddDate(-8, 0, 0)                                                       //8八年前的今天
	t := time.Date(date.Year(), date.Month(), date.Day(), 9, 30, 0, 0, date.Location()) //八年前的今天9:30
	fmt.Println("八年前的今天:", t)

	//毫秒时间戳
	fmt.Println("毫秒时间戳：", t.UnixMilli())

	//格式化
	s := t.Format("2006/01/02 15:04:05 -0700")
	fmt.Println("格式化：", s)

	//那天是周几 本年到那天过了多少周
	weekday := t.Weekday()
	_, w := t.ISOWeek()
	fmt.Printf("那天是周%d，那年过了%d周\n", int(weekday), w)

	//距离今天过了多少天
	delta := now.Sub(t)
	fmt.Printf("距离今天过了%d天", int(delta.Hours()/24))

}
