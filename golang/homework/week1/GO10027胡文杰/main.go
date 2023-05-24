package main

import (
	"fmt"
	"math/rand"
	"time"
)

//9*9乘法表
func test1() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				fmt.Printf("%d*%d=%-2d ", j, i, j*i)
			} else {
				fmt.Printf("%d*%d=%-3d ", j, i, j*i)
			}
		}
		fmt.Println()
	}
}

//9*9乘法表
func test2() {
	var width int
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				width = 2
			} else {
				width = 3
			}
			fmt.Printf("%d*%d=%-[3]*d ", j, i, width, j*i)
		}
		fmt.Println()

	}
}

//100以内20个随机数
func test3(long int) {
	var (
		sum     = 0
		product = 1
	)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < long; i++ {
		v := r.Intn(99) + 1
		if i&1 == 0 {
			sum += v
		} else {
			product *= v
		}
		//fmt.Printf("%-2d %-3d %d", v, sum, product)
		//fmt.Println()
	}
	fmt.Println(sum, product)
}

//斐波那契数列
func test4(long int) {
	a, b := 1, 1
	fmt.Println(a, b)
	for {
		a, b = b, a+b
		if b > long {
			break
		}
		fmt.Println(a, b)
	}
}

func main() {
	//test1()
	//test2()
	//test3(20)
	test4(100)

}
