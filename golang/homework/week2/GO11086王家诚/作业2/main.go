package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，
// 第偶数个累乘求积。打印结果
func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 给定一个随机数随机数生成器
	var sum0 int = 0
	var sum1 int = 1
	for i := 0; i < 20; i++ {
		a := r.Intn(99) + 1 //指定生成的范围  + 1 是为了 1-100  而不是从0开始
		if a&1 == 0 {
			sum0 = sum0 + a // a与1 相与等于0 则为奇数

		} else {
			sum1 = sum1 * a // 否则为偶数  sum1原始值为1 是为了保留第一次生成的偶数

		}

	}
	fmt.Println("奇数之和=", sum0)
	fmt.Println("偶数之积=", sum1)
}

// 批改意见
// 1. 题目要求的是奇数个和偶数个，不是奇数和偶数
