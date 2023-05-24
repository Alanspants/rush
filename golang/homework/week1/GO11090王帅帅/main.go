package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 习题1
func func1() {
	// 行数循环
	fmt.Println("习题1")
	for i := 1; i < 10; i++ {
		// 列数循环
		for j := 1; j <= i; j++ {
			fmt.Printf("%vx%v=%v\t", j, i, i*j)
		}
		fmt.Println() // 行结束换行
	}
}

// 习题2
func func2() {
	fmt.Println("习题2")
	// a. 随机20个 0<x<100
	fmt.Println("20个小于100随机数")
	rand.Seed(time.Now().UnixNano())
	lst := make([]int, 20) // 生成一个20长度切片
	for i := 0; i < 20; i++ {
		n := rand.Intn(99) + 1 // 随机包含0, 加1
		fmt.Printf("%v ", n)
		lst[i] = n // 写入每个元素
	}
	fmt.Println()
	// b. 数组奇数索引求和，偶数求积
	fmt.Println("数组奇数索引求和，偶数求积")
	var num1 int
	var num2 = 1 // 求积不能使用int默认值0
	for i := 0; i < 20; i++ {
		if (i+1)%2 != 0 { // 第多少个为索引值+1
			num1 += lst[i]
		} else {
			num2 *= lst[i]
		}
	}
	fmt.Printf("以上随机数，奇数位求和为： %v\n", num1)
	fmt.Printf("以上随机数，偶数位求积为： %v\n", num2)

}

// 习题3
func func3() {
	fmt.Println("习题3")

	// 初始三个数
	var x, y int = 0, 1
	var z int = x + y
	fmt.Printf("100以内的斐波那契数列：")

	// 在起始时输出 x,y,z
	if x == 0 && y == 1 {
		fmt.Printf("%v %v %v ", x, y, z)
	}

	for {
		// x,y,z全部向右移动一个元素,后续仅输出z元素
		x, y = y, z
		z = x + y
		if z > 100 { // 大于100退出
			break
		}
		fmt.Printf("%v ", z)
	}

}

func main() {
	func1()
	func2()
	func3()
}
