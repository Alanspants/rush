package main

import (
	"fmt"
	"math/rand"
)

func test1() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func test2() {
	sli := make([]int, 0)
	for {
		rand1 := rand.Intn(100)
		if rand1 == 0 {
			continue
		}
		sli = append(sli, rand1)
		if len(sli) == 20 {
			break
		}
	}
	fmt.Println(sli, len(sli))
	sum := 0     //定义单数初始值
	product := 1 //定义偶数初始值
	for _, v := range sli {
		if v%2 == 0 {
			product *= v
		} else {
			sum += v
		}
	}
	fmt.Printf("单数求和：%d 偶数乘积：%d\n", sum, product)
	// [81 87 47 59 81 18 25 40 56 94 11 62 89 28 74 11 45 37 6 95] 20
	// 单数求和：668 偶数乘积：2921333022720
}

func test3() {
	arr := make([]int, 0)
	arr = append(arr, 0)
	arr = append(arr, 1)
	for i := 0; i < 10; i++ {
		arr = append(arr, arr[i]+arr[i+1])
		if arr[i+2] > 100 {
			break
		}
	}
	fmt.Println(arr) //[0 1 1 2 3 5 8 13 21 34 55 89]
}

func main() {
	test1()
	test2()
	test3()
}
