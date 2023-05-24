package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("九九乘法表")
	str := ""
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				str += fmt.Sprintf("%d*%d=%-2d", j, i, i*j)
			} else {
				str += fmt.Sprintf("%d*%d=%-3d", j, i, i*j)
			}
		}
		fmt.Println(str)
		str = ""
	}

	fmt.Println("随机数")
	sum1, sum2 := 0, 1
	for i := 0; i < 20; {
		ran := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
		if ran != 0 {
			if i&^1 != i {
				sum1 += ran
			} else {
				sum2 *= ran
			}
			fmt.Println(ran)
			time.Sleep(1)
			i++
		} else {
			continue
		}
	}
	fmt.Printf("求和：%v\t乘积：%v\n", sum1, sum2)

	fmt.Println("斐波那契数列")
	num1, num2, num3 := 0, 0, 0
	for i := 1; i < 100; i++ {
		switch {
		case i == 1:
			num1 = 1
			fmt.Println(num1)
			continue
		case i == 2:
			num2 = 1
			fmt.Println(num2)
			continue
		}
		num3 = num1 + num2
		if num3 <= i {
			fmt.Println(num3)
			num1 = num2
			num2 = num3
		} else {
			continue
		}
	}
}

