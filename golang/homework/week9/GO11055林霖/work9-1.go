package main

import (
	"fmt"
	"strconv"
	"time"
)

// 时间练习
func timeTest() {

	curTime := time.Now()
	olderYear := curTime.Year() - 8
	curDay := curTime.Format("01-02")
	olderStr := strconv.Itoa(olderYear) + "-" + curDay + " 09:30:00"
	olderTime, err := time.ParseInLocation("2006-01-02 15:04:05", olderStr, curTime.Location())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("毫秒时间戳是：%d\n", olderTime.UnixMilli())
	fmt.Println(olderTime.Format("2006/01/02 15:03:04 -0700"))
	fmt.Printf("那天是 %v\n", olderTime.Weekday().String())
	days := curTime.YearDay()
	fmt.Printf("本年已经过了%d周\n", days/7)
	fmt.Printf("距今天已经过了%d天\n", int(time.Since(olderTime).Hours()/24))

}

func main() {
	// timeTest()
	// list1 := new(SingleNodeList)
	// list1.append(NewSingleNode("1")).append(NewSingleNode("2")).append(NewSingleNode("3")).append(NewSingleNode("4"))
	// list1.internodes()

	list := new(DoubleNodeList)
	list.append(NewDoubleNode("1")).append(NewDoubleNode("2")).append(NewDoubleNode("3")).append(NewDoubleNode("4"))
	list.interNodes()
	list.insert(3, NewDoubleNode("5"))
	list.insert(0, NewDoubleNode("8"))
	list.insert(6, NewDoubleNode("9"))
	list.interNodes()
	list.remove(list.size)
	list.interNodes()
	list.pop()
	list.interNodes()
	list.remove(2)
	list.interNodes()

}
