package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task1() {
	for i := 1; i <= 9; i++ {
		// fmt.Println(i)
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%[2]d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}

func task2() {
	var arr [20]int
	var sum int
	var pf int64
	rand.Seed(time.Now().UnixNano())

	// r := rand.Intn(100)
	// 非0值，不会判断
	for k := 0; k < 20; k++ {
		arr[k] = rand.Intn(100)
		// for _, v := range arr {

		// 	fmt.Println(v)
		// }

		if k%2 == 0 {
			pf *= int64(arr[k])
		} else {
			sum += arr[k]
		}
	}

	fmt.Printf("The product is:%d\n", pf)
	fmt.Printf("The sum is:%d\n", sum)
}

func task3(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return task3(n-1) + task3(n-2)
	}

}
func main() {
	task1()
	task2()
	for i := 1; i <= 100; i++ {
		if task3(i) < 100 {
			fmt.Printf("%d ", task3(i))
		} else {
			break
		}

	}
}

// 批改意见：
// 通用问题
// 1. 所有函数，养成写注释的习惯，方便自己后面查看
// 2. 变量含义尽量清晰，例如索引用idx（index的缩写），如果一开始不会可以先搜索，用的多了就熟悉了。

// 作业问题
// 1. 第二题里Intn方法返回值范围是[0, n)，因此可能会出现0值
// 2. Golang里索引也是从0开始的，因此奇数个和偶数个搞反了；
//    奇数个索引是0，2，4，8...，偶数个索引是1,3,5,7...
