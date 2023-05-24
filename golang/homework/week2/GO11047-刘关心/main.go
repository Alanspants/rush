package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("打印99乘法表")
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if (i == 3 || i == 4) && j == 2 {
				fmt.Printf("%dx%d=%d   ", j, i, j*i)

			} else {
				fmt.Printf("%dx%d=%d  ", j, i, j*i)
			}
		}
		fmt.Println()
	}

	fmt.Println("-------------------------------------")
	fmt.Println()

	fmt.Println("随机生成100以内的20个非0正整数,打印出来")

	a := rand.New(rand.NewSource(time.Now().UnixNano()))
	Sum := 0
	Mul := 1
	for b := 1; b < 22; b++ {
		Random := a.Intn(100)
		if Random == 0 {
			continue
		}

		fmt.Printf("%v ", Random)

		//单数个求和
		if b&1 == 1 {
			// fmt.Printf("|%v-->%v|\n", b, Random)
			// fmt.Printf("@%v+%v=%v--->%v@\n", Sum, Random, Sum+Random, Sum)
			Sum = Sum + Random
			// fmt.Printf("*%v*", Sum)

			//偶数个求积
		} else {
			Mul = Mul * Random

		}

	}
	fmt.Println("")
	fmt.Printf("单数个求和:%v\n偶数个求积:%v\n", Sum, Mul)

	fmt.Println("-------------------------------------")
	fmt.Println()

	fmt.Println("100以内的斐波那契数列")

	for Num1, Num2 := 1, 2; Num1 < 100; Num1, Num2 = Num2, Num1+Num2 {
		fmt.Printf("%v ", Num1)
	}

}

// 批改意见
// 1. 第二题生成的随机数是21个，不是20个。且如果返回0时跳过，那么多次返回0时，会导致实际生成的随机数少于要求的数量；
// 2. 第三题斐波那契数列是1，1，开头，你这里少了一个1，代码逻辑有问题
