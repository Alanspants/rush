package main

import (
	"fmt"
	"sort"
)

// 1、实现对圆形、三角形、长方形求面积。
// 2、利用第1题，构造3个以上图形，至少圆形、三角形、矩形各有一个，对上题的图形按照面积降序排列

// 面积area
type are struct {
	area float64
}

var pai = 3.14

// circular圆形
type Circular struct {
	are
	radius int // 半径radius
}

// 求圆形的面积
func (c *Circular) area() float64 {
	c.are.area = pai * float64(c.radius) * float64(c.radius)
	return c.are.area
}

// triangle三角形
type Triangle struct {
	are
	bottom, high int // 底,高
}

// 求三角形的面积
func (t *Triangle) area() float64 {
	t.are.area = (float64(t.bottom) * float64(t.high)) / 2
	return t.are.area
}

// 长方形rectangle
type Rectangle struct {
	are
	long, wide int // 长，宽
}

// 求长方形的面积
func (r *Rectangle) area() float64 {
	r.are.area = float64(r.long * r.wide)
	return r.are.area
}

// 实现求面积的接口，不同图形结构体结果不同
type Areaer interface {
	area() float64
}

// 多态，求各个图形的面积时统一用这个
func MakeArea(a Areaer) float64 {
	return a.area()
}

// 圆形构造函数
func NewCircular(radius int) *Circular {
	return &Circular{radius: radius}
}

// 三角形构造函数
func NewTriangle(bottom, high int) *Triangle {
	return &Triangle{bottom: bottom, high: high}
}

// 长方形构造函数
func NewRectangle(long, wide int) *Rectangle {
	return &Rectangle{long: long, wide: wide}
}

type areSlice []are

func (a areSlice) Len() int           { return len(a) }
func (a areSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a areSlice) Less(i, j int) bool { return a[i].area > a[j].area }

func main() {
	// 构建实例，一样一个
	// 图形的长宽高半径都是int整数，面积是float
	circular := NewCircular(3)
	triangle := NewTriangle(6, 7)
	rectangle := NewRectangle(6, 6)

	// 通过MakeArea求面积，每个图形结构体都有area() float64方法，实现了Areaer接口，所以都能把图形结构体实例传进去
	fmt.Println("圆形的面积:", MakeArea(circular))
	fmt.Println("三角形的面积:", MakeArea(triangle))
	fmt.Println("长方形的面积:", MakeArea(rectangle))

	// TODO:没想到 简单法把所有实例放进一个slice里。。。。。。。。。。。。。。。。。。。。
	// 构造一个存放各图形结构体实例中面积are的切片，排序用
	arer := make([]are, 0, 3)
	arer = append(arer, circular.are)
	arer = append(arer, triangle.are)
	arer = append(arer, rectangle.are)
	fmt.Println("降序排序之前的顺序：", arer)

	// 排序

	// sort.Slice(arer, func(i, j int) bool {
	// 	return arer[i].area > arer[j].area
	// 	// return arer[i].area < arer[j].area //升序
	// })

	sort.Sort(areSlice(arer))

	fmt.Println("降序排序之后的顺序：", arer)
}
