package main

import (
	"fmt"
	"math/rand"
	"time"
)

// MultiplicationTable 九九乘法表
func MultiplicationTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				fmt.Printf("%d*%d=%-2d", j, i, j*i)
			} else {
				fmt.Printf("%d*%d=%-3d", j, i, j*i)
			}
		}
		fmt.Println()
	}
}
func MultiplyTable() {
	var width int
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 1 {
				width = 2
			} else {
				width = 3
			}
			//fmt.Printf("%d*%d=%-[4]*[3]d ", j, i, i*j, width)
			fmt.Printf("%d*%d=%-[3]*d ", j, i, width, i*j)
			if j == i {
				fmt.Println()
			}
		}
	}
}

// OddSumAndEvenProduct 奇数项求和，偶数项求积
func OddSumAndEvenProduct(length int) (int, float64) {
	s := make([]int, 0, length*5) //如果是make([]int, length, length*5)将从length+1开始追加，前面length个长度值全是0,len(s)将为40
	var random int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var (
		sum     int     = 0
		product float64 = 1
	)
	for i := 0; i <= length; i++ {
		// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
		random = r.Intn(99) + 1 //Intn(99)是[0,99)左闭右开，实际是0到98(包含)
		if i&1 == 0 {
			sum += random
		} else {
			product *= float64(random)
		}
		s = append(s, random)
	}
	fmt.Println(s)
	/*	for j, e := range s {
		if (j+1)%2 == 1 {
			sum += e
		} else {
			product *= e
		}
	}*/
	return sum, product
}

// Fibonacci 斐波那契数列
func Fibonacci() {
	var arr = [100]int{0, 1} //数组给定前两个元素值是为了保证循环中arr[i] = arr[i-2] + arr[i-1]不为0
	var slice = []int{0, 1}  //循环从数组第三个位置开始，切片前两个元素要占上
	for i := 2; i < 46; i++ {
		arr[i] = arr[i-2] + arr[i-1]
		if arr[i] > 1134903170 {
			//fmt.Println(i)
			goto Label
		}
		slice = append(slice, arr[i])
	}
Label:
	fmt.Println(slice)
}

// Fbi 递归 增加cache
var m = map[uint64]uint64{0: 0, 1: 1, 2: 1}

func Fbi(i uint64) uint64 {
	if i == 0 || i == 1 {
		return i
	}
	if _, ok := m[i]; !ok {
		m[i] = Fbi(i-2) + Fbi(i-1)
	}
	return m[i]
}
func Fib(limiter int) {
	a, b := 0, 1
	fmt.Println(b)
	for {
		a, b = b, a+b
		if b > limiter {
			break
		}
		fmt.Println(b)
	}
}
func main() {
	//fmt.Printf("%T %[1]f, %t\n", math.Pow(99, 10), math.MaxUint64 > math.Pow(99, 10))
	MultiplyTable()
	//fmt.Println("-----------------------------------------------------------------------------------------------------")
	//sum, product := OddSumAndEvenProduct(20)
	//fmt.Printf("第单数个值累加求和为：%d，第双数个累乘求积为：%.2f\n", sum, product)
	Fibonacci()
	fmt.Println("==============================================================")
	Fib(1134903170)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	for i := 0; i < 46; i++ {
		fmt.Println(Fbi(uint64(i)))
	}
	// var c = "\x63" //string长度为1个字节\x表示转义x为16进制
	// fmt.Printf("%T %[1]s %d\n", c, len(c))
	// var a = 0x63
	// fmt.Printf("%T %[1]v\n", string(a)) //string(a) 把a当做Unicode兼容ASCII去查表
	// var b = 00001111                    //00开头是八进制
	// fmt.Printf("%T %[1]b %[1]d\n", b)
}
