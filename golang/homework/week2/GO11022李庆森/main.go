package main

import (
	"fmt"
	"math/rand"
	"time"
)

func multiplication() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if (i == 3 || i == 4) && (j == 3) {
				fmt.Printf(" %d*%d=%d  ", j, i, j*i)
			} else {
				fmt.Printf("%d*%d=%d  ", j, i, j*i)
			}

		}
		fmt.Println()

	}
}

func random() {
	var sum, mu int = 0, 1
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //随机数生成器
	for i := 0; i < 20; {
		num := r.Intn(100)
		if num == 0 {
			continue
		}
		fmt.Printf("%d ", num)
		if i&1 == 0 {
			sum = sum + num
		} else {
			mu = mu * num
		}
		i++
	}
	fmt.Println("")
	fmt.Printf("第单数个(不是索引)累加求和: %d\n第偶数个累乘求积: %d\n", sum, mu)
}

func fibonacci() {
	for a, b := 0, 1; a < 100; {
		fmt.Printf("%d ", a)

		a, b = b, a+b
	}
}

func main() {
	fmt.Println("打印九九乘法表。如果可以要求间隔均匀")
	multiplication()

	fmt.Println("随机生成100以内的20个非0正整数,打印出来")
	random()

	fmt.Println("打印100以内的斐波那契数列")
	fibonacci()
}

// 批改意见
// 1. 函数内记得写注释，方便自己后面看代码
// 2. 第二题里for循环逻辑错误，Intn可能会返回0，但是出现0时跳过，那么对应的数组元素就是默认值0，最终的结果就不符合要求。
// 3. 斐波那契数列是从1开始，不是从0开始
