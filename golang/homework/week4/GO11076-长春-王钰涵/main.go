package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*第一题作答
1.请问s1、s2内各有什么元素？
答：s1：[0 0 0] s2：[0 0 0 1]
2.s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？
答：s1修改元素一定会影响s2，s2只要不修改s2[4]就会影响s1
3.s2再增加一个元素会怎么样？
答：会更换底层数组，s1和s2再无关系
*/

//根据数组创建切片，切片元素是数组相邻2项的和
func SliceAdd() {
	a := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	s := make([]int, 0, 10)
	for i := 1; i < len(a); i++ {
		s = append(s, a[i]+a[i-1])
	}
	s1 := s[1:3]
	fmt.Printf("%v\n%v", s, s1)
}

//作业排序第一版（冒泡排序）
func MP() {
	a := rand.New(rand.NewSource(time.Now().UnixNano()))
	sum := 1
	var fSlice = make([]int, 0, 100) //make([]int, 0, 10)
	for i := 0; i < 100; i++ {
		fSlice = append(fSlice, a.Intn(401)-200)
	}
	for x := 0; x < len(fSlice); x++ {
		for y := x + 1; y < len(fSlice); y++ {
			if fSlice[x] < fSlice[y] {
				num := fSlice[x]
				fSlice[x] = fSlice[y]
				fSlice[y] = num
			}
		}
	}
	fmt.Println(fSlice)
	for i := 0; i < len(fSlice); i++ {
		if i != len(fSlice)-1 && fSlice[i] == fSlice[i+1] {
			sum += 1
		} else if i == len(fSlice)-1 {
			fmt.Println(fSlice[i], sum)
		} else {
			fmt.Println(fSlice[i], sum)
			sum = 1
		}
	}
}

//作业排序第二版（sort排序）
func SortPX() {
	var s = make([]int, 0, 100)
	var sum = 1
	a := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		s = append(s, a.Intn(401)-200)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && s[i] == s[i+1] {
			sum += 1
		} else if i == len(s)-1 {
			fmt.Println(s[i], sum)
		} else {
			fmt.Println(s[i], sum)
			sum = 1
		}
	}
}

func main() {
	SliceAdd()
	//MP()
	SortPX()
}
