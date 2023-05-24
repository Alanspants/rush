package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1、打印九九乘法表。如果可以要求间隔均匀。
func theFirstOne() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}

// 2、随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，
// 第偶数个累乘求积。打印结果
func randNum() {
	// rands := rand.New(rand.NewSource(time.Now().UnixMicro()))
	rands := rand.New(rand.NewSource(time.Now().UnixNano())) // 这个地方注意使用的函数
	var (
		sum  = 0
		duct = 1
	)
	for i := 0; i < 20; i++ {
		num := rands.Intn(100)
		fmt.Printf("第%d个随机数：%d\n", i, num)
		if i%2 == 0 {
			sum += num
		} else {
			duct = duct * num
		}
	}
	fmt.Printf("随机数累加求和：%d,随机数乘积：%d", sum, duct)
}

// 3、打印100以内的斐波那契数列
func feibo(n int) {
	// 1,1,2,3,5,8,13
	num := make([]int, n, n)
	for i := 0; i < n; i++ {
		if i <= 1 {
			num[i] = 1
		} else {
			num[i] = num[i-1] + num[i-2]
		}
	}
	fmt.Println(num)

}

func main() {
	theFirstOne()
	randNum()
<<<<<<< HEAD
	feibo(11)
=======
	// feibo(11)
>>>>>>> 0fe9d4cf9cc6d37cd44a2c2c5cfc4cc3baf13dbb

}
