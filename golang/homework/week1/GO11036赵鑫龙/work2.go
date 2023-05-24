package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum() {
	multiply := 1        //用于乘法计算
	summation := 0       //用于加法计算
	s := make([]int, 0)  //用来存储100以内的20个随机数
	s1 := make([]int, 0) //用来存储偶数个
	s2 := make([]int, 0) //用来存储单数个
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano()) //设定种子
		r := rand.Intn(99) + 1
		s = append(s, r)
		if i&1 == 0 {
			s1 = append(s1, r)
			summation += r
		} else {
			s2 = append(s2, r)
			multiply *= r
		}
	}
	fmt.Printf("随机生成一百以内的20个随机数:\n%d\n", s)
	fmt.Printf("单数个：%d\n", s1)
	fmt.Printf("偶数个：%d\n", s2)
	fmt.Printf("偶数个相乘结果为: %d\n", multiply)
	fmt.Printf("单数个相加结果为: %d\n", summation)
}

func main() {
	sum()
}
