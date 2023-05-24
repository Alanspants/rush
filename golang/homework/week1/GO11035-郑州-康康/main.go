//main 包注释

package main

import (
	"fmt"
	"math/rand"
)

// 九九乘法表
func multiplicationTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d*%d=%2d ", j, i, j*i)
			if i == j {
				fmt.Println()
				break
			}
		}

	}
}

// 九九乘法表2
func multiplicationTable2() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if i*j < 10 {
				fmt.Printf("%d*%d=%d  ", j, i, j*i)
			} else {
				fmt.Printf("%d*%d=%d ", j, i, j*i)

			}

		}
		fmt.Println()
	}
}

// 随机数相加、相乘
func randomMultiplication() {
	var sum int
	var m int = 1
	for i := 1; i < 21; i++ {
		var n int
		n = rand.Intn(100)
		if n == 0 {
			n = rand.Intn(100)
			fmt.Println(i, n)
			if i&1 == 1 {
				sum += n
			} else {
				m *= n
			}
			continue
		}
		fmt.Println(i, n)
		if i&1 == 1 {
			sum += n
		} else {
			m *= n
		}

	}
	fmt.Println("随机数第单数个求和：", sum)
	fmt.Println("随机数第偶数个求和：", m)

}

func fibonacciSeries() {
	var a int = 1
	var b int = 1
	// var t int
	// 1 1 2 5
	// 	3 2
	for {
		b, a = a, b
		b += a
		if b > 100 {
			break
		}
		fmt.Println(b)
	}
	//输出结果不要漏掉最开始的元素
}

func main() {
	// multiplicationTable()
	multiplicationTable2()
	randomMultiplication()
	fibonacciSeries()
}
