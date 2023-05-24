package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX = 100
	MIN = 0
)

// 生成随机生成值
func generateRandomValue() int {
	// 将初始化过程放到函数里面会导致每次生成的随机数一样
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(MAX-MIN) + MIN
	fmt.Println("v is: ", v)
	return v
}

// 计算 累积 累和
func computer(n int) (int, int) {
	sums := 0
	products := 1
	randomNumberArrary := [20]int{}
	for i := 0; i < n; i++ {
		r := generateRandomValue()
		randomNumberArrary[i] = r
		// 这个地方加和乘的值都不对
		if (i+1)%2 == 1 {
			sums += n
		} else {
			products *= n
		}

	}
	fmt.Println(randomNumberArrary)
	return sums, products
}
func main() {
	var n int = 20
	s, p := computer(n)
	fmt.Println(s, p)
}
