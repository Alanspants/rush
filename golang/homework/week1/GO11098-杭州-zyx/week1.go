package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 打印九九乘法表。如果可以要求间隔均匀。
func jiujiu() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			//fmt.Println(i + "*" + j + "="i*j)
			// 输出格式有问题，注意换行符
			fmt.Printf(strconv.Itoa(j) + "*" + strconv.Itoa(i) + "=" + strconv.Itoa(i*j) + "\t")
		}
		println()
	}
}

// 随机生成100以内的20个非0正整数，打印出来。对生成的数值，第单数个（不是索引）累加求和，
// 第偶数个累乘求积。打印结果
func randNum() {
	var arr [20][2]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i][0] = rand.Intn(100) + 1
		if i < 2 {
			arr[i][1] = arr[i][0]
		} else {
			if i%2 == 0 {
				arr[i][1] = arr[i-2][1] + arr[i][0]
			} else {
				arr[i][1] = arr[i-2][1] * arr[i][0]
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			fmt.Println("第" + strconv.Itoa(i+1) + "位的值为:" + strconv.Itoa(arr[i][0]) + ",对应的奇数位累加结果为:" + strconv.Itoa(arr[i][1]))
		} else {
			fmt.Println("第" + strconv.Itoa(i+1) + "位的值为:" + strconv.Itoa(arr[i][0]) + ",对应的偶数位累乘结果为:" + strconv.Itoa(arr[i][1]))
		}
	}
}

// 打印100以内的斐波那契数列
func feiBoNaQie() {
	var arr [101]int64
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= 100; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	for j := 0; j < len(arr); j++ {
		fmt.Println("第" + strconv.Itoa(j) + "位的值为:" + strconv.FormatInt(arr[j], 10))
	}
}

func main() {
	fmt.Println("作业1:==================打印九九乘法表=====================")
	jiujiu()
	fmt.Println()
	fmt.Println("作业2:随机生成100以内的20个非0正整数，打印出来。第单数个累加求,第偶数个累乘求积。")
	fmt.Println()
	randNum()
	fmt.Println()
	fmt.Println("作业3:打印100以内的斐波那契数列================================")
	fmt.Println()
	feiBoNaQie()
}
