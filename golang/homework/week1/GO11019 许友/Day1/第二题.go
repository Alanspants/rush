package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成20个100以内的非0整数
func RandNum() []int {
	var l []int
	rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(101))

	for {
		randn := rand.Intn(101)
		// fmt.Println(randn)
		if randn != 0 {
			l = append(l, randn)
			if len(l) == 20 {
				break
			}
		}

	}
	return l
}

func main() {
	l := RandNum()
	// fmt.Println(l)

	var odd []int
	var even []int
	var sum_odd int
	var multify int = 1

	//分离奇数和偶数存入不同分片
	for i, v := range l {
		// fmt.Println(l[i] / 2)
		// 这个地方判断奇数个和偶数个数字的逻辑错误，请重新尝试
		if i%2 == 1 {
			odd = append(odd, v)

		} else {
			even = append(even, v)
		}
	}
	fmt.Println(odd)
	fmt.Println(even)

	//单数求和
	for i := 0; i < len(odd); i++ {
		sum_odd = sum_odd + odd[i]
	}
	fmt.Printf("odd number's sum is %d\n", sum_odd)

	//偶数求积
	for i := 0; i < len(even); i++ {
		multify = multify * even[i]

	}
	fmt.Printf("multify number's sum is %d\n", multify)

}
