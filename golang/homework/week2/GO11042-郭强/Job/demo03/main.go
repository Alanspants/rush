package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SumAndPro(s []int) (sum, product int) {
	product = 1
	for _, v := range s {
		if v%2 != 0 {
			sum += v
		} else {
			product *= v
		}
	}
	return sum, product
}
func main() {
	//随机生成100以内的20个非0正整数
	rand.Seed(time.Now().UnixNano())
	var randArr []int
	for i := 0; i < 20; i++ {
		r := rand.Intn(100)
		if i == 0 {
			continue
		}
		randArr = append(randArr, r)
	}
	fmt.Println("100以内非零的随机正整数为:", randArr)

	sum, product := SumAndPro(randArr)
	fmt.Println("所有随机数中单数累计和为:", sum)
	fmt.Println("所有随机数中偶数乘积为:", product)
}

// 批改意见
// 1. 生成随机数时，如果返回0就跳过，那么最终数组内有效值数量可能会少于20个。这地方的逻辑有问题。
// 2. 题目要求奇数个求和，偶数个求乘积，不是奇数和偶数
