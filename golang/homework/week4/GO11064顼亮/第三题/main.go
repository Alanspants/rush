package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 3、数字重复统计
// 随机产生100个整数
// 数字的范围[-200, 200]
// 升序输出这些生成的数字并打印其重复的次数
func main() {
	// var s = [8]int{1, 4, 9, 16, 2, 5, 10, 15}
	// fmt.Println(s)
	rand.Seed(time.Now().Unix())
	var s [100]int
	for i := 0; i < 100; i++ {
		s[i] = rand.Intn(400) + 1 - 200

	}
	fmt.Println("初始数组: ", s)
	s1 := s[:]
	sort.Ints(s1)
	fmt.Println("升序后: ", s1)
	var ct int = 0
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1); j++ {
			if i == j {
				break
			} else {
				if s1[i] == s1[j] {
					println("重复的数字: ", s1[i])
					ct++
				}
			}
		}
	}

	println("重复的次数: ", ct)
}

// 批改意见
// 1. Intn的范围是[0, n)，因此Intn(400)+1的范围是[1, 400]，-200之后就是[-199, 200]，这块主要是对Intn的范围不熟悉
// 2. 不是要你找出重复数字，而是统计每个随机数出现的次数
