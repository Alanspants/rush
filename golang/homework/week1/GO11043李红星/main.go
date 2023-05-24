package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	MultiTable(9) //99乘法表

	odd, even, _ := RandNum(100) //100以内随机数 求和 乘积 运算
	fmt.Println(odd, even)

	fib(100)
}

// MultiTable 99乘法表
func MultiTable(a int) {
	for i := 1; i <= a; i++ { //行
		for j := 1; j <= i; j++ { //列
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}

// RandNum 100以内随机数 求和 乘积 运算
//奇数位相加 偶数位相乘
func RandNum(a int) (b int, c int64, err error) {
	if a <= 1 {
		return 0, 0, errors.New("need >1")
	}
	var (
		oddAddRet     int
		evenDivideRet = 1
	)
	rand.Seed(time.Now().UnixNano())
	var numbyte []int
	for i := 1; i < 20; i++ {
		num := rand.Intn(a)
		if num == 0 {
			i--
			continue
		}
		numbyte = append(numbyte, num)
	}
	//fmt.Println(numbyte)
	for i := 0; i < len(numbyte); i++ {
		if (i+1)%2 != 0 {
			oddAddRet += numbyte[i]
		} else {
			evenDivideRet *= numbyte[i]
		}
	}
	return oddAddRet, int64(evenDivideRet), nil
}

//100以内的斐波那契数列
func fib(a int) {
	//1、1、2、3、5、8、13、21、34、55、89
	ret := 1
	fibbyte := []int{}
	for i := 0; i < a; i++ {
		if i <= 1 {
			fibbyte = append(fibbyte, ret)
		} else {
			ret = fibbyte[i-2] + fibbyte[i-1]
			if ret > 100 {
				break
			}
			fibbyte = append(fibbyte, ret)
		}
	}
	fmt.Println(fibbyte)
}
