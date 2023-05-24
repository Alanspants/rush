package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMicro())
	sum := 0
	pro := 1
	for i := 0; i < 20; i++ {
		if i&1 == 0 {
			sum = sum + (rand.Intn(99) + 1)
		} else {
			pro = pro * (rand.Intn(99) + 1)
		}
	}
	fmt.Printf("和为:%v,积为:%v", sum, pro)
}

// 批改意见
// 1. 要注意pro默认类型是int32，最终乘积的值可能会超过它能表示的范围
