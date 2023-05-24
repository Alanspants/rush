package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1、99乘法表
func multiplicationTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				fmt.Printf("%d*%d=%-2d", j, i, j*i)
			} else {
				fmt.Printf("%d*%d=%-3d", j, i, j*i)
			}
		}
		fmt.Println()
	}
}

// 2、随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和,第偶数个累乘求积。打印结果
func randomDemo() {
	sum := make([]int, 0, 100)
	//伪随机，它是一个固定的计算公式，初始值为seed,从seed计算得到一个所谓的随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i <= 20; i++ {
		//r随机数生成器
		random := r.Intn(100)
		if random <= 0 {
			continue
		}
		sum = append(sum, random)
	}
	fmt.Println(sum)

	var add int
	var multi = 1
	for k, v := range sum {
		if k%2 == 0 {
			add += v
		} else {
			multi *= v
		}
	}
	fmt.Printf("单数个相加等于：%d\n", add)
	fmt.Printf("偶数个相乘等于：%d", multi)
}

// 作业2：随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和,第偶数个累乘求积。打印结果
func randomDemo2() {
	limiter := 20
	var (
		sum     = 0
		product = 1
	)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < limiter; i++ {
		random := r.Intn(99) + 1
		fmt.Println(random)
		if random%2 == 0 {
			sum += random
		} else {
			product *= random
		}
	}
	fmt.Printf("单数个相加等于：%d\n", sum)
	fmt.Printf("单数个相乘等于：%d\n", product)
}

// 作业三：打印100以内的斐波那契数列
func fibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 作业三：打印100以内的斐波那契数列
func fib(limiter int) {
	a, b := 1, 1
	fmt.Println(b)
	for {
		a, b = b, a+b
		if b > limiter{
			break
		}
		fmt.Println(b)
	}

}

func main() {
	//作业一：九九乘法表
	//multiplicationTable()

	//作业二：随机生成100以内的20个非0正整数，打印出来。第单数个累加求,第偶数个累乘求积
	//randomDemo()
	//randomDemo2()
	//作业三:打印100以内的斐波那契数列
	//fibs := []int{}
	//for i := 1; i <= 20; i++ {
	//	for ret := range fibs{
	//		fmt.Println(ret)
	//	}
	//	fibs = append(fibs, fibonacci(i))
	//}
	//fmt.Println(fibs)
	fib(100)

}
