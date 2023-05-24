package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 99乘法表
func MulList() {
	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {
			if x*y < 10 && y != 1 {
				fmt.Printf("%d*%d=%d  ", y, x, y*x)
			} else {
				fmt.Printf("%d*%d=%d ", y, x, y*x)
			}
		}
		fmt.Println()
	}
}

//随机20个数，然后位数为单的相加求和，位数为双的相乘求积
func RandomS(length int) {
	var aSlice = make([]int, 0, length+5)
	var (
		sum int
		mul = 1
	)
	// 生成新的伪随机种子
	a := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 切片数据赋值
	for i := 0; i < length; i++ {
		if b := a.Intn(100); b == 0 {
			aSlice = append(aSlice, b+1)
		} else {
			aSlice = append(aSlice, b)
		}
	}
	// 判断单数位还是双数位，单数相加，双数相乘
	for x, y := range aSlice {
		fmt.Printf("%d ", y)
		if x&1 == 1 {
			mul *= y
		} else {
			sum += y
		}
	}
	fmt.Println()
	fmt.Printf("位数为单号的累加和为：%d，位数为双数的累乘积为：%d", sum, mul)
}

// 斐波那契数列
func FbList() {
	var aSlice = make([]int, 0, 15)
	fmt.Println("输出一百以内的斐波那契数列")
	for i := 0; i < 99; i++ {
		// fmt.Println("i的值为", i)
		switch {
		case i == 0:
			aSlice = append(aSlice, 0)
			fmt.Printf("%d ", aSlice[i])
			continue
		case i == 1:
			aSlice = append(aSlice, 1)
			fmt.Printf("%d ", aSlice[i])
			continue
		case aSlice[i-1]+aSlice[i-2] > 100:
			// fmt.Printf("aSlice[i-1]+aSlice[i-2] = %d,跳出循环", aSlice[i-1]+aSlice[i-2])
			goto label
		default:
			// fmt.Println("aslice的len：", len(aSlice))
			aSlice = append(aSlice, aSlice[i-1]+aSlice[i-2])
			fmt.Printf("%d ", aSlice[i])
		}
	}
label:
}

func main() {
	MulList()
	RandomS(20)
	FbList()
}
