package main

import (
	"fmt"
	"math/rand"
)

func test1() {
	num := 10
	for i := 1; i < num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %-2d ", j, i, j*i)
		}
		fmt.Println()
	}
}

func test2() {
	sum1 := 0
	var sum2 uint64 = 1
	for i := 1; i <= 20; i++ {
		num := rand.Intn(100)
		for num == 0 {
			num = rand.Intn(100)
			if num != 0 {
				break
			}
		}
		if i&1 == 0 {
			sum2 *= uint64(num)
		} else {
			sum1 += num
		}
	}
	fmt.Printf("奇数位和为：%d\n", sum1)
	fmt.Printf("偶数位积为：%d\n", sum2)
}

func test3() {
	arr := make([]int, 2, 100)
	arr[0] = 1
	arr[1] = 1
	for i := 2; ; i++ {
		num := arr[i-1] + arr[i-2]
		if num > 100 {
			break
		}
		arr = append(arr, num)
	}
	fmt.Println("斐波那契数列：", arr)
}

func main() {
	test1()
	test2()
	test3()

}

// 批改意见
// 答案都完成的不错
// 函数里记得写注释，方便自己以后查看
