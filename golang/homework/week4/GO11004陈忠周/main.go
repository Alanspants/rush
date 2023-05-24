package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//1、看下面一段程序，回答问题
//s1 := make([]int, 3, 4)
//s2 := append(s1, 1)
//请问s1、s2内各有什么元素？
//s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？
//s2再增加一个元素会怎么样？
/*答：s1 [0 0 0]
      s2 [0 0 0 1]
	  s1修改会影响s2， s2元素修改会影响s1，因为它们共用同一个底层数组
	  s2再增加一个元素会发生扩容，之后就分道扬镳了
*/
func Slice() {
	s1 := make([]int, 3, 4)
	fmt.Printf("%p %p; %d %d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	s2 := append(s1, 1)
	fmt.Printf("%p %p; %d %d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	for i := 0; i < len(s1); i++ {
		s1[i] = i + 30
	}
	fmt.Println(s1)
	fmt.Println(s2)

	s2[0] = 100
	fmt.Println(s1, s2)
	fmt.Println("-----------")
	s1 = append(s1, 200)
	fmt.Println(s1, s2)
	fmt.Println("-----------")
	s2 = append(s2, 300)
	fmt.Printf("%p %p; %d %d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("%p %p; %d %d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
}

// 2、有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。
func NewSlice() {
	arr := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	fmt.Println(arr)
	fmt.Println("------------")
	count := len(arr) - 1
	out := make([]int, 0, count)
	for i := 0; i < count; i++ {
		out = append(out, arr[i]+arr[i+1])
	}
	fmt.Println(out)
}

// 3、数字重复统计
// 随机产生100个整数
// 数字的范围[-200, 200]
// 升序输出这些生成的数字并打印其重复的次数
func NumberRepeatCount() {
	//(1) 生成随机数100个，范围[-200,200)
	//当前时间的纳秒值做随机数种子：r
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var nums []int
	for i := 0; i < 100; i++ {
		nums = append(nums, r.Intn(400)-200)
	}
	fmt.Println(nums)
	//(2) 创建map：key存放nums中的随机数，value存放随机数出现的次数；
	m := make(map[int]int, 100) //容量至多为100

	//（3）将随机数(nums)作为key放入map中，value为重复次数
	// 从nums[0]开始取随机数，首先判断nums keys中是否有该随机数，没有则添加key，并value = 1; 否则， value += 1;
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok { //如果是true，则表示key中已有该值，value += 1
			m[nums[i]] += 1
		} else { //如果false，表示key中没有该值，values = 1
			m[nums[i]] = 1
		}
	}
	//（4）将key值取出存入切片中（方便排序）
	//需要将map的key值排序，需要先将key值放入切片，后对切片排序
	arr := make([]int, 0, 100)
	for k, _ := range m {
		arr = append(arr, k)
	}

	//fmt.Println("排序前：", arr, len(arr))
	//（5）排序
	sort.Ints(arr)
	//fmt.Println("排序后：", arr, len(arr))

	//（6）打印，查看随机数与重复次数
	//sum := 0
	for i := 0; i < len(arr); i++ {
		v := m[arr[i]]
		fmt.Printf("随机数%d：%d，出现次数：%d\n", i+1, arr[i], v)
		//sum += v
	}
	//fmt.Println(sum)
}

func main() {
	// Slice()
	// NewSlice()
	NumberRepeatCount()
}
