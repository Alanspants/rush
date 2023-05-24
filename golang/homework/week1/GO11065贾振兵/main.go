/*
 * @Description:
 * @Author: ZhenbingJia
 * @Date: 2022-12-26 17:35:44
 * @LastEditTime: 2022-12-29 11:20:07
 * @LastEditors: ZhenbingJia
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 99乘法表
func multiplication() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

// 100以内随机数
func random() {
	var arr []int
	sum := 0
	pro := 1
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 20; i++ {
		arr = append(arr, rand.Intn(99)+1)
	}
	fmt.Println(arr) // 作业第一步，打印100内非0整数

	for k, _ := range arr {
		if k%2 == 0 {
			sum += arr[k]
		} else {
			pro *= arr[k]
		}

	}
	fmt.Printf("第单数之和为 %d,第偶数之乘积为 %d\n", sum, pro) // 作业第2、第3步
}

// 斐波那契数列
func fibonacci(a int) int {
	if a < 2 {
		return 1
	}
	return fibonacci(a-1) + fibonacci(a-2)
}

func main() {
	multiplication() // 作业一
	random()         // 作业二

	// 作业三
	for d := 0; d <= 10; d++ {
		result := fibonacci(d)
		fmt.Println(result)
	}
}
