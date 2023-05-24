package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 作业1
// 打印九九乘法表
func multiplicationTable() {
	fmt.Println("---------Go语言打印九九乘法表----------")
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			//fmt.Println(i, j)
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
			//var re = fmt.Sprintf("%d*%d=%d", i, j, i*j)
			//fmt.Print(re + " ")
		}
		fmt.Println()
	}

}
func randomNumber() {
	fmt.Println("---------Go语言随机生成二十个非0整数，并单个数相加求和，第偶数个累乘求积----------")
	rand.Seed(time.Now().UnixNano())
	numberList := make([]int, 0, 100)
	fmt.Println()
	for i := 0; i <= 20; i++ {
		random := rand.Intn(100)
		if rand.Intn(100) <= 0 {
			continue
		}
		numberList = append(numberList, random)
	}
	fmt.Printf("生成的二十个随机数为：%v \n", numberList)

	var add int
	var multi int = 1
	for k, v := range numberList {
		//fmt.Println(k, v)
		if k%2 == 0 {
			add += v
		} else {
			multi *= v
		}
	}
	fmt.Printf("单数位置的相加等于：%d\n", add)
	fmt.Printf("偶数位置的相乘等于：%d", multi)
	fmt.Println()
}

func FibonacciSequence() {
	fmt.Println("---------打印100以内的斐波那契数列----------")
	f1 := 1
	f2 := 1
	fmt.Printf("%d ", f1)
	fmt.Printf("%d ", f2)

	for fn := 1; fn <= 100; fn++ {

		if fn == f2+f1 {
			fmt.Printf("%d ", fn)
			f1, f2 = f2, fn
		}

	}
	fmt.Println()

}
func main() {
	//作业一
	multiplicationTable()
	//作业二
	randomNumber()
	//作业三
	FibonacciSequence()

}
