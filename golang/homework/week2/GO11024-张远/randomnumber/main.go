package main

import (
	"fmt"
	"math/rand"
	"time"
)

func creatint() (result int) {
	rand.Seed(time.Now().UnixNano())
	var data int = 0
	for i := 0; ; i++ {
		data = rand.Intn(100)
		fmt.Println(data)
		if data > 0 && data%2 == 0 {
			break
		}
	}
	result = data
	return result

}
func main() {
	var data1 int = 1
	var data2 int = 0
	var randdata int
	fmt.Println("20个非0正整数")
	for i := 1; i <= 20; i++ {
		randdata = creatint()
		fmt.Println(randdata)
		if i%2 == 0 {
			data1 = data1 * randdata

		} else {
			data2 = data2 + randdata
		}

	}
	fmt.Println("偶数累乘积", data1)
	fmt.Println("单数累加和", data2)

}

// 批改意见
// 1. 生成随机数的逻辑没有看明白，按照createint函数里的逻辑，返回的只有奇数，没有偶数，还可能返回0，不符合要求
