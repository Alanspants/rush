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
	//1.毫秒时间戳是多少？
	now := time.Now()                     //计算当前时间
	eightYearAgo := now.AddDate(-8, 0, 0) //计算8年前的今天
	fmt.Println(now)
	fmt.Println(eightYearAgo)
	fmt.Println("~~~~~~~~~~~")
	eightYearsAgoDate := time.Date(eightYearAgo.Year(), eightYearAgo.Month(), eightYearAgo.Day(), 9, 10, 0, 0, eightYearAgo.Location())
	timestamp := eightYearsAgoDate.UnixNano() / int64(time.Millisecond)
	fmt.Println("8 years ago today at 9:30:00 的毫秒时间戳是：", timestamp)
    milli := eightYearsAgoDate.UnixMilli()
	fmt.Println("8 years ago today at 9:30:00 的毫秒时间戳是：", milli)

	// 2. 格式化输出时间为 2005/09/10 21:35:40 +0800 的形式
	timeStr := "2005/09/10 21:35:40 +0800"
	parse, err := time.Parse("2006/01/02 15:04:05 -0700", timeStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("~~~~~~~~~~~")
	fmt.Println(parse)
	//3. 请问那天是周几？到那天，本年已经过了多少周？
	weekday := eightYearsAgoDate.Weekday()
	fmt.Println("八年前的今天是：", weekday)
	eightYearAgoBeginning := time.Date(eightYearAgo.Year(), 1, 1, 0, 0, 0, 0, eightYearAgo.Location())
	sub := eightYearsAgoDate.Sub(eightYearAgoBeginning)
	weeks := int(sub.Hours() / 24 / 7)
	fmt.Printf("八年前的今天是:%v, 到那天，本年已经过了%d周\n", weekday, weeks)

}
