package main

import (
	"fmt"
)

// loop
func Factorialv1(n int) int { // uint64
	if n < 1 {
		panic("n不能小于1")
	}
	p := 1
	for i := 2; i <= n; i++ {
		p = p * i
	}
	return p
}

// recursion v1 优美的公式，便于理解
func Factorialv2(n int) int {
	if n < 1 {
		panic("n不能小于1")
	} else if n == 1 {
		return 1
	}
	return Factorialv2(n-1) * n
}

// recursion v2 loop 也没有什么优势，只是为了练习
func Factorialv3(n int, p int) int {
	if n < 0 {
		panic("n不能小于1")
	} else if n == 1 {
		return p
	} else if n == 0 {
		return 1
	}
	return Factorialv3(n-1, p*n)
}

func PrintTrianglev1(n int) {
	width := 40 // 定死了，右对齐的宽度怎么确定？用最后一行
	var line string
	for i := 1; i <= n; i++ { // i 控制行 i=3
		line = ""                 // 清空
		for j := i; j >= 1; j-- { // j=3;j>=1;j--  j=2;j>=1 j=1;j>=1 j=0
			// fmt.Print(j, " ") // 3 2 1
			// 拼行
			line += fmt.Sprint(j, " ")
		}
		// fmt.Printf("%[1]*[2]s\n", width, line)
		// fmt.Printf("%[1]*s\n", width, line)
		fmt.Printf("%*s\n", width, line)
		// fmt.Printf("%[2]*[1]s\n", line, width)
	}
}

func PrintTrianglev2(n int) {
	// width := 40 // 定死了，右对齐的宽度怎么确定？用最后一行
	// 能否先做最后一行？
	last := ""
	for i := n; i >= 1; i-- {
		last += fmt.Sprintf("%d ", i)
	}
	// fmt.Println(last)
	width := len(last)
	var line string
	for i := 1; i < n; i++ {
		line = "" // 清空
		for j := i; j >= 1; j-- {
			line += fmt.Sprint(j, " ") // 拼行
		}
		fmt.Printf("%*s\n", width, line)
	}
	fmt.Println(last)
}

func PrintTrianglev3(n int) {
	last := ""
	for i := n; i >= 1; i-- {
		if i == 1 {
			last += fmt.Sprint(i)
		} else {
			last += fmt.Sprintf("%d ", i)
		}
	}
	// fmt.Println(last)
	width := len(last) // max index = width -1
	for i := width - 1; i >= 0; i-- {
		if last[i] == 32 {
			fmt.Printf("%*s\n", width, last[i+1:])
		}
	}
	fmt.Print(last)
}

func main() {
	// fmt.Println(Factorialv1(6)) // 720 思考小，3
	// fmt.Println(Factorialv2(6))
	fmt.Println(Factorialv3(6, 1))
	// PrintTrianglev1(12)
	// PrintTrianglev2(12)
	PrintTrianglev3(12)
}
