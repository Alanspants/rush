package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("第一题")
	//1、看下面一段程序，回答问题
	//s1 := make([]int, 3, 4)
	//s2 := append(s1, 1)
	//
	//请问s1、s2内各有什么元素？
	//答： s1中元素 [0 0 0 ]  s2中元素[0 0 0 1]
	//s1修改一个元素会影响s2吗？
	//答：会影响。s1和s2共用一个底层数组  修改s1中[ 0 0 0 ]任意数 都会影响s2
	//s2修改一个元素会影响s1吗？
	//答：不一定 。 s2修改 切片中前三个数会影响 s1的值  ，s2修改第四个数不会影响s1的值
	//s2再增加一个元素会怎么样？
	//答：会发生扩容，内存中申请一个新的长度为8的底层数组，将旧的底层数组值赋给新的数组中
	//之后追加新增的元素值    s2的len++    因为s2和s1不再共有一个底层数组 二者之后也没有任何关系了

	fmt.Println("第二题----")
	//2、有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。
	var arr = [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	var slice []int
	//遍历arr数组 将数组相邻结果追加到切片中
	for i := 0; i < len(arr); i = i + 2 {
		slice = append(slice, arr[i]+arr[i+1])
	}
	fmt.Println(slice) //[5 25 7 25]

	fmt.Println("第三题------")
	//3、数字重复统计
	//随机产生100个整数
	//数字的范围[-200, 200]
	//升序输出这些生成的数字并打印其重复的次数

	var slice2 []int //声明一个slice切片 用来保存100个随机数
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		slice2 = append(slice2, rand.Intn(401)-200)
	}

	sort.Ints(slice2)
	fmt.Println("排序后", slice2)
	fmt.Println("长度=", len(slice2))

	num := make(map[int]int, 100)
	for _, v := range slice2 {
		num[v] += 1
	}
	fmt.Println("切片重复数是", num)
}

// 批改意见
// 1. 第二题里让相邻的两个数相加，即1和2，2和3，3和4相加，不只是1和2，3和4
