// 数据结构测试题
package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

//切片说明
func slicesPrint() {
	// TODO
	s1 := make([]int, 3, 4) //len3,cap 4 == 字面量定义的 []int{0,0,0}
	s2 := append(s1, 1)     //len4,cap 4 == 字面量定义的 []int{0,0,0,1}
	fmt.Println(s1, s2)
	//没有扩容公用底层数组，s1,s2,修改均会影响
	//s2在添加元素后，扩容了。这个时候生成新的底层数组，分道扬镳，s1和s2独立了，在没有联系
}

// TODO 第二次作业
// 有一个数组 [1,4,9,16,2,5,10,15]，生成一个新切片，要求新切片元素是数组相邻2项的和。
// a := [...]int{1, 4, 9, 16, 2, 5, 10, 15}
func arrySum() {
	a := [8]int{1, 4, 9, 16, 2, 5, 10, 15}
	s := make([]int, 7)
	for i := 1; i < len(a); i++ {
		s[i-1] = a[i] + a[i-1]
	}
	fmt.Println(a)
	fmt.Println(s, len(s), cap(s))
}

//冒泡排序
// 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 针对所有的元素重复以上的步骤，除了最后一个。
// 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
func bubbleSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j+1] < a[j] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func bubbleSort2(a []int) []int {
	if len(a) < 2 {
		return a
	}
	n := len(a)
	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
	return a
}

//选择排序
// 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置。
// 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
func selectionSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	n := len(a)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
	return a
}

//插入排序
// 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
// 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
func insertionSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	n := len(a)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0 && a[j] > a[j+1]; j-- {
			a[j], a[j+1] = a[j+1], a[j]
		}
	}
	return a
}

func insertionSort2(a []int) []int {
	if len(a) < 2 {
		return a
	}
	n := len(a)
	for i := 1; i < n; i++ {
		newNumIndex := i
		for {
			if newNumIndex-1 >= 0 && a[newNumIndex-1] > a[newNumIndex] {
				a[newNumIndex-1], a[newNumIndex] = a[newNumIndex], a[newNumIndex-1]
				newNumIndex--
			} else {
				break
			}
		}
	}
	return a
}

// TODO 作业
// 	随机产生100个整数
// 数字的范围[-200, 200]  rand.Intn(400)-200
// 升序输出这些生成的数字并打印其重复的次数
func randCount(n int) {
	rand.Seed(time.Now().UnixMicro())
	s := make([]int, n)
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		r := rand.Intn(399) - 200
		s[i] = r
		// if _, ok := m[r]; !ok {
		// 	m[r] = 1
		// } else {
		// 	m[r] += 1
		// }
		// m[r] += 1
		m[r]++
	}
	fmt.Println(s)
	// s = bubbleSort(s)
	// s = bubbleSort2(s)
	// s = selectionSort(s)
	// s = insertionSort(s)
	s = insertionSort2(s)
	// sort.Ints(s)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(s)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// for key, v := range m {
	// 	fmt.Printf("key:%-4d-----count:%d\n", key, v)
	// }
}

func showSortFun() {
	var (
		maxLen   = 5
		maxValue = 100
		testTime = 1
	)
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < testTime; i++ {
		s := lenRandSlice(maxLen, maxValue)
		s1 := copySlices(s)
		fmt.Println(s1)
		s2 := selectionSort(s)
		if !isSorted(s2) {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("排序错误了-----------------")
		}
		fmt.Println(s2)

	}
}

//是否有序
func isSorted(s []int) bool {
	min := 0
	for i := 1; i < len(s); i++ {
		if s[min] > s[i] {
			return false
		} else {
			min = i
		}
	}
	return true
}

func copySlices(s []int) []int {
	s2 := make([]int, len(s))
	copy(s2, s)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("%p,%p,%v,%v\n", &s, &s2, s, s2)
	// for i := 0; i < len(s); i++ {
	// 	s2[i] = s[i]
	// }
	return s2
}

//得到随机切片,对数器
func lenRandSlice(maxLen, maxValue int) []int {
	len := rand.Intn(maxLen) + 1
	s := make([]int, len)
	for i := 0; i < len; i++ {
		s[i] = rand.Intn(maxValue)
	}
	return s

}

// TODO 3-4作业
// 求n的阶乘。至少使用递归函数完成一次
func factorial(n int) uint64 {
	if n <= 1 {
		return 1
	}
	var cur uint64 = 1
	for i := 1; i <= n; i++ {
		cur *= uint64(i)
	}
	return cur
}

//递归版本
func factoria2(n int) uint64 {
	if n <= 1 {
		return 1
	}
	return uint64(n) * factoria2(n-1)
}

//递归改循环
func factoria3(n int, b uint64) uint64 {
	if n <= 1 {
		return b
	}
	return factoria3(n-1, uint64(n)*uint64(n-1))
}

var stack = list.New()

func factoriaStack(n int) int {
	if n <= 1 {
		return 1
	}
	stack.PushBack(1)
	for i := 2; i <= n; i++ {
		a := stack.Back()
		stack.Remove(a)
		stack.PushBack(a.Value.(int) * i)
	}
	a := stack.Back()
	result := stack.Remove(a)
	return result.(int)

}

//有点像插入排序，算法有问题，还得思考一下
func printTable(n int) {
	//外面1-n打印行数
	for i := 0; i < n; i++ {
		//打1-n-1 空格
		for m := 0; m < n; m++ {
			g := n - 1 - i
			length := 0
			if m < g {
				if g <= 3 {
					length = 3 * g
				} else {
					length = 2*g + 3
				}
				for h := 0; h < length; h++ {
					fmt.Print("\x20")
				}
				break
			}
		}
		//打印数字
		for k := 0; k < n; k++ {
			if k < n-1-i {
			} else {
				if n-k >= 10 {
					fmt.Printf("%-3d", n-k)
				} else {
					fmt.Printf("%-2d", n-k)
				}
			}
		}
		fmt.Println()

	}

}

func printTest() {
	// n := 11
	for i := 0; i < 11; i++ {
		// length := 0
		// if i < 10 {
		// 	length = 2 * (9 - i)
		// } else {
		// 	length = 2 * (9 - i)
		// }
		for j := 0; j < 2*(10-i); j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++ {
			// fmt.Print(i - k)
			if i < 10 {
				fmt.Printf("%-2d", i-k)
			} else {
				fmt.Printf("%-3d", i-k)
			}

		}
		fmt.Println()
	}
}

func printTest2(n int) {
	//打印空格
	//0-n-1
	//0-n-2
	for end := n - 1; end >= 0; end-- {
		//打印“X”
		// for i := 0; i < end; i++ {
		// 	fmt.Print("x")
		// }
		for j := 1; j <= n; j++ {
			if j > end {
				// if n = 10 {
				// 	fmt.Printf("%-3d", n-j+1)
				// } else {
				// 	fmt.Printf("%-2d", n-j+1)
				// }
				// fmt.Printf("%-3d", n-j+1)
				print(n - j + 1)
			}
		}
		fmt.Println()
	}
}

func main() {
	// fmt.Println(factoriaStack(3))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// fmt.Println(factoria2(3))
	randCount(100)
}

// 批改意见
// 1. s2修改最后一个元素不会影响s1
// 2. Intn的范围是[0, n)，因此Intn(399)-200会漏掉2个值，199和200
// 3. 代码中记得写注释，方便后面自己查看
