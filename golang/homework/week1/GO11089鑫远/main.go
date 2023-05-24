package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber()
	// 打印乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println("")
	}

}

// 随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，第偶数个累乘求积。打印结果
func randomNumber() {
	var cunt int
	var product int
	for i := 1; i <= 20; i++ {
		num := rand.Intn(100)
		// 需要排除非0正整数
		fmt.Println(num)
		if i%2 == 0 {
			product = product * num
		} else if i%2 != 0 {
			cunt = cunt + num
		}
	}
	fmt.Println("第单数个累加求和", cunt)
	fmt.Println("累乘求积", product)
}
