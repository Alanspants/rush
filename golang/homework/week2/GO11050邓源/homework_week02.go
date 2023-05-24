package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1. 九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			result := i * j
			// 注意一下i和j的顺序
			if result < 10 && j == 1 {
				fmt.Printf("%d*%d=%d ", i, j, result)
			} else if result < 10 {
				fmt.Printf("%d*%d=%d  ", i, j, result)
			} else {
				fmt.Printf("%d*%d=%d ", i, j, result)
			}
		}
		fmt.Print("\n")
	}

	// 2. 随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，第偶数个累乘求积
	fmt.Println("--------------------------")
	rand.Seed(time.Now().UnixNano())
	var num [20]int
	sum := 0
	quad := 1
	for i := 0; i < 20; i++ {
		num[i] = rand.Intn(99) + 1
		if i%2 == 0 {
			sum += num[i]
		} else if i%2 == 1 {
			quad *= num[i]
		}
	}
	fmt.Printf("随机数为：%v \n第单数个求和为: %d \n第偶数个求积为: %d \n", num, sum, quad)

	// 3. 斐波那契数列
	fmt.Println("--------------------------")
	fmt.Print("斐波那契数: ")
	var fib [3]int
	fib[0] = 0
	fib[1] = 1
	for i := 0; i < 100; i++ {
		if i == 0 || i == 1 {
			fmt.Printf("%v ", fib[i])
			continue
		}
		fib[2] = fib[0] + fib[1]
		if fib[2] > 100 {
			break
		}
		fmt.Printf("%v ", fib[2])
		fib[0] = fib[1]
		fib[1] = fib[2]
	}
}
