package main

import (
	"fmt"
	"time"
)

//第一题时间操作
func TimeZY() {
	t := time.Now()
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	date := t.AddDate(-8, 0, 0)
	t2 := time.Date(date.Year(), date.Month(), date.Day(), 9, 30, 00, 00, location)
	fmt.Printf("八年前的时间为：%v\n", t2)
	fmt.Printf("八年前的时间戳为(毫秒)：%v\n", t2.UnixMilli())
	now := time.Now()
	sub := now.Sub(t2)
	fmt.Printf("8年前的那个时间是周%d，8年前的时间距今有%.1f周\n", date.Weekday(), sub.Hours()/(7*24))
	fmt.Printf("距离今天已经过了%.1f天\n", sub.Hours()/24)
}

func main() {
	TimeZY()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}
