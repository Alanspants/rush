package main

import (
	"fmt"
	"math/rand"
	"time"
)

func multiplyTable1() {
	for i := 1; i < 10; i++ {
		// 做了9行，i表示某一行。每一行里面还要做9列
		for j := 1; j <= i; j++ {
			if j == 1 {
				fmt.Printf("%d*%d=%-2d", j, i, i*j)
			} else {
				fmt.Printf("%d*%d=%-3d", j, i, i*j)
			}
		}
		fmt.Println()
	}
}

func multiplyTable2() {
	var width int // 零值
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				width = 2
			} else {
				width = 3
			}
			// fmt.Printf("%d*%d=%-[4]*[3]d", j, i, i*j, width) // %[4]d
			fmt.Printf("%d*%d=%-[3]*d", j, i, width, i*j) // %[4]d
		}
		fmt.Println()
	}
}

func multiplyTable3() {
	var width int // 零值
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				width = 2
			} else {
				width = 3
			}
			// fmt.Printf("%d*%d=%-[4]*[3]d", j, i, i*j, width) // %[4]d
			fmt.Printf("%d*%d=%-[3]*d", j, i, width, i*j) // %[4]d
			if i == j {
				fmt.Println()
			}
		}
		// fmt.Println()
	}
}

func calc() {
	limiter := 20
	var ( // 把业务上有关联的可以批量定义在一起
		sum     = 0
		product = 1
	)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < limiter; i++ {
		v := r.Intn(99) + 1 // [0, 99) 100以内不含0 [1, 99]
		fmt.Println(v)
		if i&1 == 0 { // i % 2 == 0
			sum += v
		} else {
			product *= v
		}
	}
	fmt.Println(sum, product)
}

func calc1() {
	limiter := 20
	var ( // 把业务上有关联的可以批量定义在一起
		sum            = 0
		product uint64 = 1
	)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < limiter; i++ {
		v := r.Intn(99) + 1 // [0, 99) 100以内不含0 [1, 99]
		fmt.Println(v)
		if i&1 == 0 { // i % 2 == 0
			sum += v
		} else {
			product *= uint64(v) // v:int, product:uint64
		}
	}
	fmt.Println(sum, product)
}

// 斐波那契数列是面试常考问题，素数问题也是常问问题
func fib() {
	// 给出最早2项，必须给
	a, b := 1, 1
	fmt.Println(a, b)
	for {
		a, b = b, a+b
		if b > 100 {
			break
		} // 退出条件
		fmt.Println(b)
	}
}

func fib1(limiter int) {
	// var limiter = 100
	// 给出最早2项，必须给
	a, b := 1, 1
	fmt.Println(a)
	for {
		a, b = b, a+b
		if a > limiter {
			break
		} // 退出条件
		fmt.Println(a)
	}
}

// func main() {
// 	// main函数必须在main包中
// 	// 同一个目录，包名必须一样
// 	// main是程序入口，操作系统调用时，必须知道从哪里开始
// 	fib1(88) // 通用代码
// 	fib1(100)
// }
