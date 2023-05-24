package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//题目一
	for i := 1; i < 10; i++ {
		for o := 1; o < 10; o++ {
			if o > i {
				continue
			}
			fmt.Printf("%v * %v = %v\t", o, i, o*i)
		}
		fmt.Printf("\n")
	}

	//题目二

	n := 0
	c := 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i < 21; i++ {
		tmp := r.Intn(100)
		if i%2 == 0 {
			if tmp != 0 {
				c = c * tmp
			}

		} else {
			n = n + tmp
		}
	}
	fmt.Println(n)
	fmt.Println(c)

	//题目三

	num1 := 1
	num2 := 1

	for i := 1; i < 100; i++ {
		if i < 3 {
			fmt.Printf("%v\t", num1)
		} else {
			i = num1 + num2
			num1 = num2
			num2 = i

			if i >= 100 {
				break

			}
			fmt.Printf("%v\t", i)
		}
	}

}

// 批改意见
// 1. 对奇数个和偶数个的判断逻辑有问题，搞反了。
// 2. 当Intn返回值为0时的特殊情况没有予以判断，跳过了。不符合要求
