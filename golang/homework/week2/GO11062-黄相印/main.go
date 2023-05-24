package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1、打印九九乘法表。如果可以要求间隔均匀。
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	//2、随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，
	//第偶数个累乘求积。打印结果
	var arr = make([]int, 0)
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 20; i++ {
		//src := rand.New(rand.NewSource(time.Now().UnixNano()))
		r100 := src.Intn(100)
		if r100 <= 0 {
			i -= 1
			continue
		}
		arr = append(arr, r100)
	}
	fmt.Println(arr)
	var sum, mul = 0, 1
	for i, v := range arr {
		if i%2 == 0 {
			mul *= v

		} else {
			sum += v
		}
	}
	fmt.Printf("单数位求和:%d 偶数位求积:%d\n", sum, mul)

	//3、打印100以内的斐波那契数列
	for a, b := 0, 1; ; {
		c := a + b
		a, b = b, c
		if c < 100 {
			fmt.Println(c)
			//continue
		} else {
			break
		}
	}

}

// 批改意见
// 1. 第二题中对arr进行迭代取值时，返回的索引i是从0开始的，因此i%2==0时是奇数，此时应该求和，而不是求乘积
// 2. 斐波那契数列前两个数是1，1，需要特殊处理，你的代码返回只有1个，逻辑有问题
