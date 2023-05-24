package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 九九乘法表
	fmt.Println("1.九九乘法表")
	work1()
	//随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，第偶数个累乘求积。打印结果
	fmt.Println("2.随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，第偶数个累乘求积。打印结果")
	work2()
	// 打印100以内的斐波那契数列
	fmt.Println("3.打印100以内的斐波那契数列")
	work3()
	fmt.Println("3.打印100以内的斐波那契数列for循环")
	work3_1()
}

func work1() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-2d ", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}

func work2() {

	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	s1 := []int{}
	s2 := []int{}
	s3 := []int{}
	s2Sum := 0
	s3Total := 0
	for i := 0; i < 20; i++ {
		num := myRand.Intn(98) + 1
		s1 = append(s1, num)
		// 这里的逻辑有问题，最终的奇数个数字和偶数个数字的数量不一样。
		if num%2 == 0 {
			s2 = append(s2, num)
			s2Sum += num
		} else {
			s3 = append(s3, num)
			if s3Total == 0 {
				s3Total = 1 * num
			} else {
				s3Total = s3Total * num
			}
		}

	}
	fmt.Printf("随机数列 %v \n", s1)
	fmt.Printf("单数求和%v为%d \n", s2, s2Sum)
	fmt.Printf("偶数求乘%v为%d \n", s3, s3Total)
}

func work3() {
	s1 := []int{}

	for i := 1; ; i++ {
		num := fn1(i)
		if num < 100 {
			s1 = append(s1, num)
		} else {
			break
		}

	}
	fmt.Println(s1)
}

func fn1(num int) int {
	if num <= 2 {
		return 1
	} else {
		return fn1(num-1) + fn1(num-2)
	}
}

func work3_1() {
	max := 100
	now := 1
	bf := 1
	bff := 1
	for i := 1; ; i++ {
		if i <= 2 {
			now = 1
		} else {
			now = bf + bff
		}
		if now > max {
			break
		}

		fmt.Printf("%d ", now)
		bff = bf
		bf = now
	}
}
