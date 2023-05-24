package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SliceOp 1、看下面一段程序，回答问题
//s1 := make([]int, 3, 4)
//s2 := append(s1, 1)
//请问s1、s2内各有什么元素？ s1:[0 0 0], s2:[0 0 0 1]
//s1修改一个元素会影响s2吗？ 会
//s2修改一个元素会影响s1吗？ 会
//s2再增加一个元素会怎么样？ s1长度容量均不变，生成s3指针地址发生了变化, 长度加1，容量翻倍扩容
func SliceOp() {

	s1 := make([]int, 3, 4)
	s2 := append(s1, 1)
	fmt.Printf("s1指针地址:%p, s2指针地址:%p\n", s1, s2)
	//请问s1、s2内各有什么元素？ s1:[0 0 0], s2:[0 0 0 1]
	fmt.Printf("s1:%v, s2:%v\n", s1, s2)
	//s1修改一个元素会影响s2吗？ 会
	s1[0] = 33
	fmt.Printf("修改s1后：  s1:%v, s2:%v\n", s1, s2)
	//s2修改一个元素会影响s1吗？ 会
	s2[1] = 34
	fmt.Printf("修改s2后：  s1:%v, s2:%v\n", s1, s2)
	s2[3] = 4
	fmt.Printf("修改s2最后一个元素后：  s1:%v, s2:%v\n", s1, s2)

	//s2再增加一个元素会怎么样？
	//s1,s2长度容量均不变
	//新增一个元素后生成s3，s3指针地址发生了变化, 长度加1，容量翻倍扩容
	s3 := append(s2, 3)
	fmt.Printf("s1长度：%d, 容量:%d; s2长度：%d, 容量:%d; s3长度：%d, 容量:%d\n", len(s1), cap(s1), len(s2), cap(s2), len(s3), cap(s3))
	fmt.Printf("s1:%v, s2:%v, s3:%v\n", s1, s2, s3)
	fmt.Printf("s1指针地址:%p, s2指针地址:%p, s3指针地址:%p\n", s1, s2, s3)
	//修改s3影响s1,s2吗？ 不影响
	s3[2] = 41
	fmt.Printf("s1:%v, s2:%v, s3:%v\n", s1, s2, s3)
	fmt.Printf("s1指针地址:%p, s2指针地址:%p, s3指针地址:%p\n", s1, s2, s3)
	fmt.Println()
}

// NewArray 2、有一个数组 [1, 4, 9, 16, 2, 5, 10, 15]，
//生成一个新切片，要求新切片元素是数组相邻2项的和。
func NewArray(arr []int) []int {
	newArr := make([]int, 0, len(arr)-1)
	for idx, ele := range arr {
		if idx < len(arr)-1 {
			newArr = append(newArr, ele+arr[idx+1])
		}
	}
	return newArr
}

// GenerateRandIntArr 3、数字重复统计
//随机产生100个整数
//数字的范围[-200, 200]
func GenerateRandIntArr(min, max int) []int {
	arr := make([]int, 0, 100)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		arr = append(arr, r.Intn(max-min)+min)
	}
	return arr
}

// AscArrAndCount 升序输出这些生成的数字并打印其重复的次数
func AscArrAndCount() ([]int, map[int]int) {
	arr := GenerateRandIntArr(-200, 200)
	fmt.Println("生成的原始数组", arr, "\n")
	num := len(arr)
	for i := 0; i < num; i++ {
		for j := 0; j < num-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	arrMap := make(map[int]int, len(arr))
	for _, ele := range arr {
		arrMap[ele]++
	}
	return arr, arrMap
}
func main() {
	SliceOp()

	arr := []int{1, 4, 9, 16, 2, 5, 10, 15, 56, 200}
	aa := NewArray(arr)
	fmt.Println("新切片，新切片元素是数组相邻2项的和:", aa)
	fmt.Println()

	ascArr, arrCount := AscArrAndCount()
	fmt.Println("升序的数组：", ascArr)
	fmt.Println()
	fmt.Println("生成的数字及其重复的次数：", arrCount)
}

// 批改意见
// 1. s2修改最后一个元素不会影响s1
// 2. Intn的范围是[0, n)，因此Intn(400)-200会漏掉最后一个随机数200
