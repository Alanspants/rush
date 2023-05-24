golang11/homework/week9/GO11035-李永康
package main

import (
	"fmt"
	"time"
)

func Time() {
	t := time.Now()
	// s := t.Format("2006/01/02 15:04:05 -0700")
	s1 := t.AddDate(-8, 0, 0)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t2 := time.Date(s1.Year(), s1.Month(), s1.Day(), 9, 30, 0, 0, loc)
	fmt.Println(t2)
	fmt.Printf("t2.UnixMilli(): %v\n", t2.UnixMilli())
	fmt.Printf("t2.Format(\"2006/01/02 15:04:05 -0700\"): %v\n", t2.Format("2006/01/02 15:04:05 -0700"))
	year, week := t2.ISOWeek()
	fmt.Println(year, "年第几周：", week)
	fmt.Printf("星期%v\n", int(t2.Weekday()))
	fmt.Printf("本年过去第%v天\n", t2.YearDay())

}
