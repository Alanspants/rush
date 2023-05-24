// 这是包注册
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//印九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%[2]d*%[1]d=%-4[3]d", i, j, i*j)
		}
		fmt.Println()
	}

	fmt.Println("#########################")
	//随机生成100以内的20个非0正整数
	sum := 0   //定义求和默认
	multi := 1 //定义求积默认
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		a := r.Intn(100)
		if a == 0 {
			i -= 1
		}
		fmt.Printf("%-3d", a)
		if a != 0 && a%2 == 0 {
			multi *= a
		} else {
			sum += a
		}
	}
	fmt.Println()
	fmt.Println("单数求和是:", sum)
	fmt.Println("双数求积是:", multi)

	fmt.Println("#########################")
	//打印100以内的斐波那契数列
	var (
		b int
		c int
		d int = 1
	)
	for i := 1; ; i++ {
		if i == 1 {
			b = 1
		} else if i == 2 {
			b = 1
		} else {
			c = b
			b = d + c
			d = c
		}

		if b < 100 {
			fmt.Printf("%-3d", b)
		} else {
			break
		}
	}

}

// 批改意见
// 1. 第二题，题目要求第奇数个元素求和，偶数个元素求乘积。不是奇数和偶数。
// 2. 函数内记得写注释，方便自己以后看代码
