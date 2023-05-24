package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Mcl()
	RandNum()
	//FbSeq()
}

// Mcl  打印九九乘法表
func Mcl() {
	//第一层循环i表示行数，总共九行
	for i := 1; i <= 9; i++ {
		//第二层循环j表示列数，列数小于等于行数
		for j := 1; j <= i; j++ {
			//格式化拼接
			fmt.Printf("%v*%v=%v\t", j, i, i*j)
		}
		//换行
		fmt.Println()
	}
}

// RandNum 随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，
// 第偶数个累乘求积
func RandNum() {
	var a int
	var b int64 = 1
	randNums := make([]int, 20)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		randNums[i] = rand.Intn(99) + 1
		//20个非0随机正整数
		fmt.Printf("%v\t", randNums[i])
	}
	fmt.Println() //换行
	// 这个地方是要求对第奇数个，偶数个数字进行处理。不是奇数、偶数进行处理
	for i, v := range randNums {
		if v%2 != 0 {
			a += randNums[i]
		} else {
			b = b * int64(randNums[i])
		}
	}
	fmt.Printf("奇数之和为:%v\n", a)
	fmt.Printf("偶数之积为:%v\n", b)
}

// FbSeq 打印100以内的斐波那契数列
func FbSeq() {
	var a int = 1
	var b int = 1
	fmt.Printf("%v\t", a)
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\t", b)
		num := a
		a = b
		b = num + a
	}
	fmt.Println()
}
