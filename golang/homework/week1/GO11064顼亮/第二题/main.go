package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main入口
func main() {
	rand.Seed(time.Now().Unix())
	var s [20]int
	for i := 0; i < 20; i++ {
		r := rand.Intn(99)
		s[i] = r + 1
	}
	fmt.Println("随机生成100以内的20个非0正整数:", s)

	var sum int
	var p uint64
	sum = 0
	p = 1
	for i := 0; i < 20; i++ {
		if i%2 == 1 {
			sum = sum + s[i]
		} else {
			p = p * uint64(s[i])
		}

	}
	fmt.Println("奇数求和: ", sum)
	fmt.Println("偶数乘积: ", p)
}
