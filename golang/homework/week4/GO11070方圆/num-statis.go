package main

import (
	"fmt"
	"math/rand"
	"time"
)

//数字重复统计
//随机产生100个整数
//升序输出这些生成的数字并打印其重复次数

func main() {
	//种子
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	sli := make([]int, 0, 1000)
	for i := 0; i < 100; i++ {
		res := r1.Intn(400) - 200
		sli = append(sli, res)
	}
	//fmt.Println(sli, len(sli))
	//升序排列
	for i := 100; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
	fmt.Println(sli, len(sli))

	for i := 0; i < 99; i++ {
		sum := 1
		if sli[i] == sli[i+1] {
			sum++
			fmt.Println(sli[i], sum)
		} else {
			fmt.Println(sli[i])
		}
	}
}

//批改意见
// 1. Intn的范围是[0, n)，因此Intn(400)-200会漏掉最后一个随机数200
