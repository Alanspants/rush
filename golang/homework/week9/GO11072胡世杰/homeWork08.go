package main

import (
	"fmt"
	"time"
)

//1、8年前的今天上午9点30分
//要求：
//1. 毫秒时间戳是多少？
//2. 格式化输出时间为 2005/09/10 21:35:40 +0800 的形式
//3. 请问那天是周几？到那天，本年已经过了多少周？
//4. 距离今天已经过了多少天了？
func get8YearsAgo() {
	now1 := time.Now()                            //获取当前时间 年月日 时分秒
	fmt.Println("当前时间:", now1)                    //2023-02-27 17:37:58
	nowDayStr := now1.Format("2006/01/02")        //对当前time解析 只要年月日
	fmt.Println("当前时间的年月日:", nowDayStr)           //2023-02-27
	strTime := fmt.Sprint(nowDayStr, " 09:30:00") //将当前年月日 和 9.30分字符串拼接
	fmt.Println("字符串 当前日期的9:30分:", strTime)
	//解析时间
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("set loadLocation err=", err)
		return
	} //加载时区
	now, err := time.ParseInLocation("2006/01/02 15:04:05", strTime, location)
	if err != nil {
		fmt.Println("ParseInLocation err=", err)
	}

	fmt.Println("当前时间:", now.Format("2006/01/02 15:04:05 -0700"))
	yearsAgo := now.Add(-time.Hour * 24 * 365 * 8)
	fmt.Println("8年前的今天:", yearsAgo.Format("2006/01/02 15:04:05 -0700"))
	fmt.Println("1970年1月1号0:00:00到8年前经过的时间戳毫秒数=", yearsAgo.UnixMilli())
	fmt.Printf("8年前的今天是%v\n", yearsAgo.Weekday())
	agoDay := now.Sub(yearsAgo).Seconds() / (60 * 60 * 24) //今天的9点30距8年前过了多少天
	fmt.Printf("到那天,本年已经过了%.2f周\n", agoDay/7)
	fmt.Printf("本年已经过了%v周,共%v天\n", now.YearDay()/7, now.YearDay())
	fmt.Printf("8年前的今天到现在已经过了%v天\n", agoDay)
}

//链表 用结构体和方法实现LinkedList链表
func main() {

	get8YearsAgo()

}
