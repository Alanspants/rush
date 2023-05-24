package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	s1 := make([]int, 3, 4)
	s2 := append(s1, 1)
	fmt.Println(s1, s2)

	s1[0] = 1
	s2[1] = 3
	fmt.Println(s1, s2)
	s2 = append(s2, 1)
	fmt.Println(s1, s2, cap(s2))

	// s1 中的元素 [0,0,0]   s2 中的元素 [0,0,0,1]

	// s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？  答案是 会，因为底层共用一个数组

	// s2再增加一个元素会怎么样？ 会扩容，假设增加一个元素 此时的 cap 容量会变为8

	// 有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。

	s3 := [...]int{1, 4, 9, 16, 2, 5, 10, 15}

	s4 := make([]int, len(s3)-1)

	for i := 0; i < len(s3)-1; i++ {
		val := s3[i] + s3[i+1]
		s4[i] = val
	}

	fmt.Println(s4) //output  [5 13 25 18 7 15 25]

	fmt.Println("--------------------------------")

	var nums []int
	numCount := make(map[int]int)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		nums = append(nums, r.Intn(401)-200)
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			numCount[nums[i]]++
		} else {
			numCount[nums[i]] = 1
		}
	}
	//书册打印值
	fmt.Println(nums)
	//统计个数
	fmt.Println(numCount)
}

// 批改意见
// 1. s2修改最后一个元素不会影响s1
// 2. 第4周了，可以尝试一下使用函数来实现上面的功能
