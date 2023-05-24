package main

import (
	"fmt"
	"strings"
)

// 习题1，求n阶乘，递归函数
func fac(n int) int {
	if n == 1 {
		return 1
	}
	return n * fac(n-1)
}

// 习题2， 接受n，打印一个倒三角（右对齐）
func tri(n int) {
	// 得到最后一行数据，以便有每行宽度w
	var w int
	var b strings.Builder
	for i := 1; i <= n; i++ {
		b.WriteString(
			fmt.Sprintf("%d ", i),
		)
	}
	w = len(strings.TrimSpace(
		b.String(),
	))

	// 每行循环输出
	for i := 1; i <= n; i++ {
		// a. 循环拼接此行字符串
		var c strings.Builder
		for j := 1; j <= i; j++ {
			// 由循环拼接，产生一行
			c.WriteString(fmt.Sprintf("%d ", i-j+1)) // 每行逆序输出算法为i-j+1
			// fmt.Printf("%[2]*[1]d ", i-j+1, w)
		}
		// b. 对此行字符串格式化输出
		fmt.Printf("%s\n", // 4. 换行
			fmt.Sprintf( // 3. 使用右对齐方式，并传入宽度w
				"%[2]*[1]s",
				strings.TrimSpace( // 2. 对此行去除收尾空格
					c.String(), // 1. 生成一行字符串
				),
				w,
			),
		)

	}

}

func main() {
	// 习题1
	fmt.Println(fac(5))

	// 习题2
	tri(12)
}
