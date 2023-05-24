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
	//时区
	tz, _ := time.LoadLocation("Asia/Shanghai")
	//  8 年前的今天
	day := time.Now().AddDate(-8, 0, 0).Format("2006/01/02")
	// 解析时间格式
	f := "2006/01/02 15:04:05"
	//解析时间 ， 8年前今天的9点30分
	t, err := time.ParseInLocation(f, fmt.Sprintf("%s 9:30:00", day), tz)
	if err != nil {
		panic(err)
	}
	fmt.Println("1、毫秒时间戳: ", t.UnixMilli())
	fmt.Println("2、输出格式: ", t.Format("2006/01/02 15:04:05 -0700"))
	_, w := t.ISOWeek()
	fmt.Printf("3、那天周%d  , 那天本年过了%d周\n", int(t.Weekday()), w)
	// d := time.Now().Sub(t)
	// fmt.Println(int(d.Hours() / 24))
	d := time.Since(t)
	fmt.Println("4、距离至今多少天: ", int(d.Hours()/24))
}
