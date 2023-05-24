package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 1、打印九九乘法表。
func multiplyTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				fmt.Printf("%d*%d=%d ", j, i, i*j)
			} else {
				fmt.Printf("%d*%d=%-2d ", j, i, i*j)
			}
		}
		fmt.Println()
	}
}

func multiplyTable1() {
	var width int
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				width = 2
			} else {
				width = 3
			}
			fmt.Printf("%d*%d=%-[3]*d", j, i, width, i*j)
		}
		fmt.Println()
	}
}

// 2、随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，
// 第偶数个累乘求积。打印结果
func printRandInt() {
	rand.Seed(time.Now().UnixNano())
	arr := [20]uint64{}
	for i := range arr {
		//arr[i] = rand.Intn(99) + 1
		arr[i] = uint64(rand.Intn(99) + 1)
		//arr[i] = 99
	}
	fmt.Println(arr)

	//flag := 1
	var sum uint64 = 0
	for i := 0; i < 20; i += 2 {
		sum += arr[i]
	}

	var mul uint64 = 1
	var index int = 1
	for index < 20 {
		if mul*arr[index] > math.MaxUint64 {
			break
		} else {
			mul *= arr[index]
			index += 2

		}
	}

	fmt.Printf("第单数个累加和为：%d\n", sum)
	if index == 21 {
		fmt.Printf("第偶数个累乘积为：%d\n", mul)
	} else {
		fmt.Println("第偶数个累乘乘积超过int64最大正整数值，无法正确保存该值！")
	}

}

// 3、打印100以内的斐波那契数列
func fbn() {
	s := make([]int, 20)
	s[0] = 1
	s[1] = 1
	l := 0
	for i := 2; i < 20; i++ {
		s[i] = s[i-1] + s[i-2]
		if s[i] > 100 {
			l = i
			break
		}
	}
	s1 := s[:l]
	fmt.Println(s1)

}

func main() {
	multiplyTable()
	//multiplyTable1()
	printRandInt()
	fbn()

}
