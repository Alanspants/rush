//8年前的今天上午9点30分
package  main

import (
    "fmt"
    "time"
)

func main() {
    //tz, _ := time.LoadLocation("Asia/Shanghai")
    var  tn = time.Now()
    fmt.Printf("毫秒时间戳为: %v\n", tn.UnixMilli())
    fmt.Println(tn.Format("2006/01/02 15:04:05 -0700"))
    t,err := time.Parse(
    "2006/01/02 15:04:05 -0700",
    "2015/02/28 09:30:00 +0800",
)
    if err != nil {
        panic(err)
    }
    fmt.Printf("8年前的今天 %v\n",t.Weekday())
    fmt.Println(t.ISOWeek())
// 3. 请问那天是周几？到那天，本年已经过了多少周？
// 4. 距离今天已经过了多少天了？
    delta := tn.Sub(t)
    fmt.Printf("delta: %T %[1]v\n",delta)
    fmt.Printf("相差天数: %d\n",int(delta.Hours()/24))
}
