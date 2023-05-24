package timecalc

import (
	"fmt"
	"math"
	"time"
)

func TimeOps() {
	timeStr := "2015-03-19 09:30:00 +0800"
	// 返回值格式，默认就是2015-03-19 09:30:00 +0800 CST格式
	t1, err := time.Parse("2006-01-02 15:04:05 -0700", timeStr)
	if err != nil {
		fmt.Println("### ", err)
	}
	// 返回值有两个，年份和本年第几周，年份值舍去
	_, w := t1.ISOWeek()
	// 根据时间差求出相差的小时数，除以24算出过去天，然后向下取整，返回值float64类型
	dayPassed := math.Floor(time.Now().Sub(t1).Hours() / 24)

	fmt.Printf("millisecond timestamp: %d\n", t1.UnixMicro())
	//fmt.Println(t1) // 2015-03-19 09:30:00 +0800 CST
	// 通过Format函数格式化为想要的时间字符串格式
	fmt.Printf("time format: %s\n", t1.Format("2006/01/02 15:04:05.9"))
	// t1.Weekday返回数字周几，通过String方法得到英文周几，因为实现了Stringer接口，可以自定义fmt.Print输出内容
	fmt.Printf("the day of that year was: %s\n", t1.Weekday().String())
	fmt.Printf("%d weeks have been paased that year\n", w)
	// 强制转换dayPassed为int类型
	fmt.Printf("%d days have passed since today\n", int(dayPassed))
}
