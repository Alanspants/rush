package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1、打印九九乘法表。如果可以要求间隔均匀
func one() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j > i {
				break
			}
			// 输出的字体间距可以再调整一下
			fmt.Printf("%2d * %2d = %2d ", j, i, i*j)
		}
		fmt.Println()
	}
}

// 随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，
// 第偶数个累乘求积。打印结果
func two() {
	sum := 0
	sum2 := 1
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 21; i++ {
		num := rand.Intn(100)
		fmt.Printf("第%d次随机数：%d\n", i, num)
		if i%2 == 0 {
			sum2 *= num
		} else {
			sum += num
		}
	}
	fmt.Printf("随机数第单数个求和: %d ,第偶数个求乘积: %d\n", sum, sum2)
}

// 打印100以内的斐波那契数列
func three() {
	a := 1
	b := 1
	sum := 1
	for i := 1; i < 100; i++ {
		if i == 1 || i == 2 {
			fmt.Println(sum)
			continue
		}
		sum = a + b // a = 1 b= 1 sum =2
		if sum > 100 {
			break
		}
		fmt.Println(sum)
		a = b
		b = sum
	}
}
func main() {
	fmt.Println("============ 第一题 ==============")
	one()
	fmt.Println("============ 第二题 ==============")
	two()
	fmt.Println("============ 第三题 ==============")
	three()
}
