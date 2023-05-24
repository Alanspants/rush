package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//99
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Printf("\n")
	}
	//随机数
	arr := [20]int{1, 2, 3}
	for j := 0; j < 20; j++ {
		r := rand.Intn(100)
		if r == 0 {
			j--
			continue
		}
		arr[j] = r

	}
	fmt.Println("---------------------------------------------------")
	var acc uint64 = 1
	var sum int16 = 0
	fmt.Printf("随机数是:")
	for i, v := range arr {
		fmt.Printf("%+v,", v)
		if (i+1)%2 == 0 {
			acc = acc * uint64(v)
		} else {
			sum = sum + int16(v)
		}

	}
	fmt.Printf("\n乘积=%d\n和=%d\n", acc, sum)
	fmt.Println("---------------------------------------------------")
	var x int = 1
	var y int = 1
	for {
		if x < 100 {
			fmt.Println(x)
			x, y = y, x+y
		} else {
			break
		}

	}

}

// 批改意见
// 1. 代码逻辑非常清晰，完成的很好
// 2. 代码内添加适量注释，可以方便自己以后看代码
