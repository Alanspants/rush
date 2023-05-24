// 1、看下面一段程序，回答问题

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// func main() {
// 	s1 := make([]int, 3, 4)
// 	s2 := append(s1, 1)
// 	fmt.Println(s1, s2)
// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~")

// 	s1[1] = 100
// 	fmt.Println(s1, s2)
// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~")

// }

/*
请问s1、s2内各有什么元素？
答:s1 [0 0 0] len=3 cap=4   s2 [0 0 0 1] len=4 cap=4

s1修改一个元素会影响s2吗？s2修改一个元素会影响s1吗？
答: s1修改会影响s2,因为s2是在s1后面追加元素,由于s1 len=3,cap=4, s2.append后s1 len=4,cap=4,没有引起扩容,公用底层数组
答: s2和共享s1底层数组,你变我就变

s2再增加一个元素会怎么样？
答: s2增加1个元素,撑爆s1底层数组,会产生新的底层数组扩容,
1.产生成新的底层数组,复制s2到新数组,append增加元素
2.s2的header指向新数组,s1和s2各有各的底层数组
*/
// // 2、有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。

func work2() {
	arr := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
	fmt.Println(len(arr))
	c1 := make([]int, 0, len(arr)) //5,13

	for i1 := 0; i1 < len(arr)-1; i1++ { //0  1	2
		c1 = append(c1, arr[i1]+arr[i1+1]) //1+4, 9+6

	}
	fmt.Println(c1)
}

// 3、数字重复统计
// 随机产生100个整数
// 数字的范围[-200, 200]
// 升序输出这些生成的数字并打印其重复的次数

func work3() {

	//1.定义1个切片
	c := make([]int, 0, 100)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//2.循环把生成的随机数,append到c切片
	for i := 0; i < cap(c); i++ {
		c = append(c, (r.Intn(400) - 200)) // 0,400 --> -200,200 //把400减去200,替换称0---> -200,200
	}

	//3.对c切片进行排序
	sort.Ints(c)
	// fmt.Println(c, len(c))

	//4.
	m := make((map[int]int), len(c))
	//通过遍历c切片，给map m《键值对》赋值,切片的c[-197]key不存在,value=1,如果c[-197]key存在,value就累加1,实现计数
	for _, v := range c { //map[-197:1 -192:2 -184:1 ]
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	// fmt.Printf("map m的顺序%v\n", m) //map m的顺序[-197:1 -192:2 -184:1 ]

	//5.切片保存 上面map的key
	c1 := make([]int, 0, 100)
	for k := range m {
		c1 = append(c1, k)
		// fmt.Println(k)
	}

	//6. 对c1排序
	sort.Ints(c1)
	// fmt.Printf("c1切片的顺序%v", c1) //c1切片的顺序[-197  -192 -184]

	for _, v := range c1 {
		// map m的顺序[-197:1 -192:2 -184:1 ]
		// c1切片的顺序[-197  -192 -184]
		fmt.Printf("随机数%v,出现%v次\n", v, m[v]) //通过遍历c1切片元素，查找map ，m[-197] , m[-192]对应的值
	}

}

func main() {
	work2()
	work3()
}

// 批改意见
// 1. 第一题，s2扩容后，s2长度比s1大1，因此s2修改最后一个元素时，s1不受影响
// 2. 第三题，Intn返回的数值范围是[0, n)，不包含最后一个元素
