package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，
//第偶数个累乘求积。打印结果
func main() {
	//种子
	sum := 0
	multi := 1
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Printf("20个随机数如下:  \n")
	for i := 0; i < 20; i++ {
		res := r1.Intn(100)
		if i%2 == 0 {
			sum += res
		} else {
			multi *= res
		}
		//如果随机数为0,则重新循环一次
		if res == 0 {
			fmt.Printf("\n")
			i--
			continue
		}
		fmt.Printf("%d ", res)
	}
	fmt.Printf("\n")
	fmt.Printf("单数和为 %d\n", sum)
	fmt.Printf("偶数之积为 %d\n", multi)
}
