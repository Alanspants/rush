package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1、打印九九乘法表。如果可以要求间隔均匀
	fmt.Println("1、打印九九乘法表。如果可以要求间隔均匀")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-3d", j, i, i*j)
		}
		fmt.Println()
	}

	// 2、随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，第偶数个累乘求积。打印结果
	fmt.Println()
	fmt.Println("2、随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，第偶数个累乘求积。打印结果")
	rand.Seed(time.Now().UnixNano())

	// 生成随机数
	var nums [20]int
	for i := 0; i < 20; i++ {
		nums[i] = rand.Intn(99) + 1
	}
	fmt.Println()
	fmt.Println("随机数为：", nums)

	// 第单数个累加求和,第偶数个累乘求积
	sum := 0
	product := 1

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			sum += nums[i]
		} else {
			product *= nums[i]
		}
	}
	fmt.Printf("累加和：%d\n", sum)
	fmt.Printf("累乘积：%d\n", product)

	// 3、打印100以内的斐波那契数列
	// F(0) = 0
	// F(1) = 1
	// F(n) = F(n-1) + F(n-2) (n ≥ 2)

	fmt.Println()
	fmt.Println("3、打印100以内的斐波那契数列")
	fib1 := 0
	fib2 := 1
	fmt.Printf("100以内的斐波那契数列为：%d, %d", fib1, fib2) // 打印输出前两项
	for {
		fibonacci := fib1 + fib2

		if fibonacci >= 100 {
			break
		}
		fib1 = fib2
		fib2 = fibonacci
		fmt.Printf(", %d", fibonacci)
	}
}
