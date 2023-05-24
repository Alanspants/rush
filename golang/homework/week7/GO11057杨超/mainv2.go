package main

import (
	"fmt"
	"math"
	"sort"
)

//实现对圆形,三角形长方形求面积
//1定义接口
type Interfase interface {
	Area() float32
}

//2.定义圆形结构体
type Circle struct {
	r float32
}

//3,定义构造函数,方便调用此函数,给结构体赋值
func NewCircle(r float32) *Circle {
	return &Circle{r: r}
}

//4.定计算面积的方法
func (c *Circle) Area() float32 {
	return math.Pi * c.r * c.r
}

type Triangle struct {
	base, height float32
}

func NewTriangle(base, height float32) *Triangle {
	return &Triangle{base: base, height: height}
}

func (t *Triangle) Area() float32 {
	return t.base * t.height / 2
}

type Rectangle struct {
	width, height float32
}

func NewRectangle(width, height float32) *Rectangle {
	return &Rectangle{width: width, height: height}
}
func (r *Rectangle) Area() float32 {
	return r.width * r.height
}
func main() {
	//这代码是简单的完成了作业要求
	c := NewCircle(4)
	fmt.Println(c, c.Area())
	t := NewTriangle(2, 3)
	fmt.Println(t, t.Area())
	r := NewRectangle(4, 5)
	fmt.Println(r, r.Area())

	// //排序
	// var i Interfase = c
	// fmt.Println(i == c) //true i就是c
	// i = t

	// //定义1个Interfase接口类型的切片,切片里的c,t,r要实现Interfase接口
	// var shapes = []Interfase{c, t, r}
	// fmt.Printf("%T", shapes)
sort.Slice(x any, less func(i, j int) bool)

}
