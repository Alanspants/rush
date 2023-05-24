package main

import "fmt"

const (
	a = "abc"
	b = 100
	c = true
)

func main1() {
	test()
}

func test() {
	const a = iota
	const b = iota
	fmt.Println(a, b)
	const ( //批量生成时候
		c = iota
		d
		e
		_ //可以用來作为标识符，但是不可以使用(空白标识符 匿名变量)
		_
		f
		g = iota
		h
		i = iota + 10
		j
		k = 30
		l //智能，Go语言，可以重复上一行的公式
		m = iota
	)
	fmt.Println(c, d, e, f, g, h, i, j, k, l, m)

	var a1 int
	a1 = 100
	fmt.Println(a1)

	var c1 = 200
	fmt.Println(c1)

	x, y, z := func() (int, int, int) {
		return 1, 2, 3
	}()
	fmt.Println(x, y, z)

}
