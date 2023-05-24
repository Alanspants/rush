package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func test4_2() {
	nums := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	slic := make([]int, len(nums)-1)
	for i, v := range nums {
		if i+1 == len(nums) {
			break
		}
		slic[i] = v + nums[i+1]
	}
	fmt.Println(slic)
}

func test4_3() {
	numMap := make(map[int]int, 100)
	keys := make([]int, 0, len(numMap))
	for i := 0; i < 100; i++ {
		num := rand.Intn(401) - 200
		keys = append(keys, num)
		v, exit := numMap[num]
		if !exit {
			numMap[num] = 1
		} else {
			numMap[num] = v + 1
		}
	}

	sort.Ints(keys)
	for _, v := range keys {
		fmt.Printf("生成的随机数：%d,重复的次数：%d \n", v, numMap[v])
	}
}

func main() {
	// s1 := make([]int, 3, 4)
	// s2 := append(s1, 1)

	//因为s1,s2共用同一个数组，所以修改s1里的元素，s2会影响，同时修改s2，s1也会影响，只不够s1只有3个元素，s2有四个元素
	//当s2再添加一个元素时，s2超过了底层数组定义的最大长度4，所以要重新创建新的数组。

	test4_2()
	test4_3()
}

// 批改意见
// 1. s2的长度比s1大，因此修改s2最后一个元素，s1不受影响
// 2. 函数中记得写注释，方便自己以后查看
