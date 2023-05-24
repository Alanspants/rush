package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
	//随机生成100以内的20个数字，第单数个求和，双数个求积
	sum := 0
	k := 1
	var ran int
	for i := 0; i < 20; i++ {
		ran = rand.Intn(99) + 1
		fmt.Println(ran)
		// 下面判断奇数个和偶数个的逻辑有问题，请重新尝试
		if i&1 == 0 { //偶数
			fmt.Println(ran)
			sum += ran
		} else {
			k *= ran
		}
	}
	v := fmt.Sprintf("累加和为：%d\n累乘乘积为：%d", sum, k)
	fmt.Println(v)

	//打印100以内斐波那契数列
	var j = 1
	sum1 := 1
	for i := 0; i < 100; {
		fmt.Printf("%d\t", sum1)
		sum1 = i + j
		i = j
		j = sum1
	}
	// 结果超出范围了
}
