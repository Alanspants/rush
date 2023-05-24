package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，第偶数个累乘求积。打印结果
func main() {
	rand.Seed(time.Now().UnixNano())
	var r []int
	sum := 0
	pruduct := 1
	for i := 0; i < 20; i++ {
		r = append(r, rand.Intn(100-1)+1)
	}
	fmt.Println(r)
	for _, v := range r {
		// 这个地方对单数个和偶数个元素逻辑错误，请检查后重新写。
		if v&1 == 0 { //这个判断逻辑在下面的输出结果是对的呀
			pruduct *= v
			// fmt.Printf("偶数：%d\n", v)
		} else {
			sum += v
			// fmt.Printf("奇数：%d\n", v)
		}

	}
	fmt.Printf("偶数乘积是：%d\n奇数和是：%d", pruduct, sum)
}
