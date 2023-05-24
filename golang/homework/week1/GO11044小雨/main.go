package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 九九乘法表1
func f1() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			if j == 1 {
				fmt.Printf("%d*%d=%-2d", j, i, i*j)
			} else {
				fmt.Printf("%d*%d=%-3d", j, i, i*j)
			}
		}
		fmt.Println()
	}
}

// 九九乘法表2
func f2() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			var width int
			if j == 1 {
				width = 2
			} else {
				width = 3
			}
			fmt.Printf("%d*%d=%-[4]*[3]d", j, i, i*j, width)
		}
		fmt.Println()
	}
}

// 九九乘法表3
func f3() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			var width int
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

// 100以内20个非0随机数，奇数个相加、偶数个相乘
func f4() {
	sum := 0
	product := 1
	v := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		v = r.Intn(99) + 1
		if i&1 == 0 {
			sum += v
		} else {
			product *= v
		}
	}
	fmt.Println(sum, product)
}

// 100以内20个非0随机数，奇数个相加、偶数个相乘
// 为保证不溢出，可将product声明为uint64
func f5() {
	sum := 0
	var product uint64 = 1
	v := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		v = r.Intn(99) + 1
		if i&1 == 0 {
			sum += v
		} else {
			product *= uint64(v) // 注意：需要强制类型转换
		}
	}
	fmt.Println(sum, product)
}

// 斐波那契
func f6() {
	a := 1
	b := 1

	fmt.Printf("%d ", a)
	for b < 100 {
		fmt.Printf("%d ", b)
		a, b = b, a+b
	}
}

func main() {
	f6()
}
