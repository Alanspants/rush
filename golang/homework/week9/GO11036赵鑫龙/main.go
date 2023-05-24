package main

import (
	"fmt"
	"time"
	"week10/linkTable"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Shanghai")                         //设置时间时区为上海
	t := time.Now().AddDate(-8, 0, 0)                                    //当前时间年份-8，得到8年前的时间
	oldTime := time.Date(t.Year(), t.Month(), t.Day(), 9, 30, 0, 0, loc) //格式化时间，将时间设置为8年前的上午9:30分
	fmt.Println("8年前的今天", oldTime)                                       //打印8年前今天上午的9:30分
	//p, err := time.Parse("2006-01-02 15:04:05 -0700 CST", fmt.Sprintf("%v", oldTime))
	fmt.Printf("毫秒时间戳: %d\n", t.UnixMilli())
	t1 := t.Format("2006/01/02 15:04:05 -0700") //以2006/01/02 15:04:05 -0700格式化当前时间
	fmt.Println(t1)
	fmt.Println(t.Weekday(), int(t.Weekday())) //分别以日期和数字形式打印今天是周几
	fmt.Println(t.ISOWeek())                   //到本日已经过去了多少天
	t2 := time.Since(t)
	fmt.Println(int(t2.Hours() / 24)) //过去的天数

	fmt.Println("----------SingleLinkList-------------")

	sl := linkTable.New()
	sl.Append(1)
	sl.Append(2).Append(3).Append(4)
	sl.IterNodes()
	fmt.Println(sl.Len())
	fmt.Println("----------DLinkList-----------------	")
	dl := &linkTable.DLinkList{}
	dl.Append(100)
	fmt.Println(dl.Len())
	//dl.Insert(4, 100)
	//dl.Insert(3, 400)
	fmt.Println(dl.IterNodes(true))
	dl.Insert(0, 300)
	fmt.Println(dl.IterNodes(true))
	dl.Insert(3, 400)
	fmt.Println(dl.IterNodes(true))
	fmt.Println(dl.Len())
	dl.Remove(1)
	fmt.Println(dl.IterNodes(true))
	fmt.Println(dl.Len())

}
