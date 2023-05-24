package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，
// 第偶数个累乘求积。打印结果

func main() {
	var sum int
	index1 := 0
	var mul int = 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		n := r.Intn(100)
		if n%2 == 0 && n != 0 {
			mul *= n
			index1++
			fmt.Println(index1, n)

		} else {
			if n == 0 {
				continue
			} else {
				sum += n
				index1++

				fmt.Println(index1, n)
			}
		}
	}
	fmt.Printf("单数和: %[1]v 偶数积: %[2]v", sum, mul)
}
