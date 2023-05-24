package main

import (
	"fmt"
	"math/rand"
)

// Multiplication 99乘法表
func Multiplication() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

// ParityEvaluation
// 随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，
// 第偶数个累乘求积。打印结果
func ParityEvaluation(n int) (int, int, error) {
	var arr = make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		for {
			a := rand.Intn(n)
			if a != 0 {
				arr = append(arr, a)
				break
			}
		}
	}
	fmt.Println(n, "以内20个非0正整数:", arr)
	var quad = 1
	var sum int
	for idx, ele := range arr {
		if (idx+1)%2 == 0 {
			quad *= ele
		} else {
			sum += ele
		}
	}
	return sum, quad, nil

}

// F 斐波那契数列
func F(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return F(n-1) + F(n-2)
}

// FibonacciWithinN N以内的斐波那契数列值
func FibonacciWithinN(n int) {
	fmt.Printf("%d%s", n, "以内的斐波那契数列值:")
	var num int
	for {
		if F(num) > n {
			break
		}
		fmt.Printf("%d ", F(num))
		num++
	}
}

func main() {
	Multiplication()
	fmt.Println()

	sum, quad, err := ParityEvaluation(100)
	if err != nil {
		panic(err)
	}
	fmt.Println("奇数求和:", sum)
	fmt.Println("偶数求积:", quad)
	fmt.Println()

	FibonacciWithinN(100)

}
