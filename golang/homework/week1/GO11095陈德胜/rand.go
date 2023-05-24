package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SumAndMultiply(s []int) (sum int, product int64) {
	product = 1
	for _, k := range s {
		if k%2 != 0 {
			sum += k
		} else {
			product *= int64(k)
			// fmt.Println(k)
		}
	}
	return sum, product
}

func main() {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	//随机生成100以内的20个非0正整数, 打印出来
	var randlist []int
	for i := 1; i <= 20; i++ {
		r := rand.Intn(100)
		if r == 0 {
			i--
			continue
		}
		randlist = append(randlist, r)
	}
	fmt.Println("随机生成数为:", randlist)

	sum, product := SumAndMultiply(randlist)
	fmt.Printf("所有随机数中的单数累计和为 %d ,", sum)
	fmt.Printf("所有随机数中的偶数的乘积为 %d 。", product)
}
