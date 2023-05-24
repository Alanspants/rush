package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("第一题")
	MulTab()
	fmt.Println("第二题")
	RandHend()
	fmt.Println("第三题")
	FiboArr()

}

func MulTab() {
	var i float32 = 1
	var j float32 = 1
	for i = 1; i < 10; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%.0f*%.0f=%-3.0f", j, i, j*i)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func RandHend() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	Mul := 1
	Sum := 0
	for i := 1; i <= 20; i++ {
		if i&1 == 0 {
			Num1 := r.Intn(98) + 1
			fmt.Printf("%d, %d\n", i, Num1)
			Mul = Mul * Num1
		} else if i&1 == 1 {
			Num2 := r.Intn(98) + 1
			fmt.Printf("%d, %d\n", i, Num2)
			Sum = Sum + Num2
		}
	}
	fmt.Printf("第单数个累加求和=%d\n", Sum)
	fmt.Printf("第偶数个累成求积=%d\n", Mul)
	fmt.Println("")
}

func FiboArr() {
	var a int = 0
	var b int = 1
	for {
		fmt.Printf("%d\n", b)
		a, b = b, a+b
		if b >= 100 {
			break
		}
	}
}

// 批改意见
// 1. Intn的返回值是[0, n)，不包含最后一个n的值，因此100以内的随机数是Intn(99)+1
// 2. 函数内记得写注释，方便自己以后看代码
