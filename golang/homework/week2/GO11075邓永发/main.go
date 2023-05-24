package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	homeWork1()
	homeWork2()
	homeWork3(100)
}

// 1. 打印 99 乘法表。如果可以要求间隔均匀。
func homeWork1() {
	for i := 1; i < 10; i++ {
		for k := 1; k <= i; k++ {
			if k == 1 {
				fmt.Printf("%d*%d=%-2d", k, i, i*k)
			} else {
				fmt.Printf("%d*%d=%-3d", k, i, i*k)
			}
		}
		fmt.Println()
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

// func homeWork1_2() {
// 	var width int
// 	for i := 1; i < 10; i++ {
// 		for k := 1; k <= i; k++ {
// 			if k == 1 {
// 				width = 2
// 			} else {
// 				width = 3
// 			}
// 			fmt.Printf("%d*%d=%-[4]*[3]d", k, i, i*k, width)
// 		}
// 		fmt.Println()
// 	}
// }

// 2. 随机生成 100 以内的 20 个非 0正整数，打印出来。
// 对生成的数值，第单数个（不是索引）累加求和，第偶数个累乘求积。打印结果
func homeWork2() {
	var (
		sum            = 0
		product uint64 = 1
	)
	src := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Printf("随机数为：")
	for i := 0; i < 20; i++ {
		k := src.Intn(99) + 1

		fmt.Printf("%v ", k)

		if i&1 == 0 {
			sum += k
		} else {
			product *= uint64(k)
		}
	}

	fmt.Println()
	// 输出第单数个累加之和
	fmt.Printf("单数个的累加之和为：%d\n", sum)
	// 输出第偶数个累乘之积
	fmt.Printf("偶数个累乘之积为：%d\n", product)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

// 3. 打印 100 以内的斐波那契数列
func homeWork3(limiter int) {
	num1, num2 := 1, 1
	fmt.Println(num1)

	for {
		num1, num2 = num2, num1+num2
		if num1 > limiter { // 退出条件
			break
		}
		fmt.Printf("%v ", num1)
	}
}
