package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1、99乘法
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Printf("\n")
	}

	// 2、随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，第偶数个累乘求积
	min := 1
	max := 100
	sum := 0
	result := 1
	for i := 1; i <= 20; i++ {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(max-min) + min
		fmt.Println(num)
		if i%2 == 0 {
			result *= num
		} else {
			sum += num
		}
	}
	fmt.Println(sum, result)

	// 3. 斐波那契数列
	num1 := 1
	num2 := 1
	fmt.Println(num1)
	fmt.Println(num2)

	for num := 0; num < 100; {
		num = num1 + num2
		if num < 100 {
			fmt.Println(num)
			num1 = num2
			num2 = num
		} else {
			break
		}

	}
}

// 批改意见
// 1. 第二题随机数初始化时，只需要初始化一次即可，多次初始化会出问题
// 2. 函数内代码记得写注释
