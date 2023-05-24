package main

//作业一
/*
问答1 s1内的元素[0 0 0]  ,s2内的元素[0 0 0 1]
问答2 s1修改元素 会影响s2，s2修改元素会影响s1.因为没扩容前共用一个底层数组，只是返回的新切片长度不一样
问答3 s2再增加一个元素会扩容。s2底层数组会变化。和s1就分道扬镳了
*/
import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//作业二
func add() {
	a := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	s := make([]int, len(a)-1)
	fmt.Printf("%p, %p,%d,%d\n", &s, &s[0], len(s), cap(s))
	for i, v := range a {
		if i >= len(a)-1 {
			break
		}
		v += a[i+1]
		s[i] = v
		//	s = append(s, v) //需要改造，目前不知道怎么改造
		fmt.Printf("%p, %p\n", &s, &s[0])
	}
	fmt.Println(s)
}

//看一下应该怎么改造
// func add1() {
// 	a := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
// 	s := make([]int, len(a)-1)
// 	fmt.Printf("%p, %p\n", &s, &s[0])
// 	for i := 0; i < len(a); i++ {
// 		if i >= len(a)-1 {
// 			break
// 		}
// 		v := a[i] + a[i+1]
// 		s = append(s, v)
// 		fmt.Printf("%p, %p,%d,%d\n", &s, &s[0], len(s), cap(s))
// 	}
// 	fmt.Println(s)
// }

//作业3
func rStatistics() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	rr := rand.New(r)
	n := make([]int, 100)
	m := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		r := (rr.Intn(200-(-200)) + (-200))
		n[i] = r
		//		n = append(n, r)
	}
	sort.Ints(n)
	// n = []int{-186, -186, -160, -150, -138, -118, -118, -108, -81, -81, 57, 79, 80, 90, 126, 142, 152, 168, 170, 170}
	// fmt.Println(n)

	for i, v := range n {
		if i >= 99 {
			break
		}
		if v == n[i+1] {
			m[v] += 1
		} else if m[v] >= 2 {
			m[n[i+1]] = 1
		} else {
			m[v] = 1
			m[n[i+1]] = 1
		}
	}
	fmt.Println(m, len(m))
}

func main() {
	//作业二
	add()
	//add1()
	//作业三
	rStatistics()

}

// 批改意见
// 1. s2修改最后一个元素不会影响s1
// 2. Intn的范围是[0, n)，因此Intn(400)-200会漏掉最后一个值200
// 3. 代码中记得写注释，方便后面自己查看
