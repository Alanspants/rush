package main

import (
	"fmt"
	"math"
	"sort"
)

//习题1：实现对圆形、三角形、长方形求面积
//1.分别构建结构体：cicle（圆形）、triangle（三角形）、rectangle（长方形）
//2.分别定义其属性：cicle: 半径; triangle: 底、高; rectangle: 长、宽
//3.定义面积方法：area()，并根据各自求面积公式不同分别实现。

//习题2：定义接口shape，拥有方法area() float64，（cicle、triangle、rectangle都已经实现了该接口）
//1.构建[]shape{}切片，切片成员为cicle、triangle、rectangle的实例
//2.实现排序

type shape interface {
	area() float64
}

// cicle结构体
type cicle struct {
	radius float64 //半径
}

// cicle构造函数
func newCicle(r float64) *cicle {
	c := &cicle{radius: r}
	return c
}

// cicle方法: 求面积
func (c *cicle) area() float64 {
	//return math.Pi * math.Sqrt(c.radius)
	return math.Pi * math.Pow(c.radius, 2)
}

//---------------------------------------------------

type triangle struct {
	bottom float64
	height float64
}

// triangle构造方法
func newTriangle(bottom float64, height float64) *triangle {
	t := &triangle{bottom: bottom, height: height}
	return t
}

// triangle方法
func (t *triangle) area() float64 {
	return t.height * t.bottom / 2
}

// ---------------------------------------------------

type rectangle struct {
	length float64
	width  float64
}

// rectangle构造方法
func newRectangle(length float64, width float64) *rectangle {
	r := &rectangle{length: length, width: width}
	return r
}

// rectangle的方法
func (r *rectangle) area() float64 {
	return r.length * r.width
}

//习题2：

func main() {
	//--------圆形求面积:cicle------------
	c := newCicle(3)
	fmt.Println(c.area())

	//--------三角形求面积:triangle-------
	t := newTriangle(3, 4)
	fmt.Println(t.area())

	//--------长方形求面积:rectangle------
	r := newRectangle(10, 8)
	fmt.Println(r.area())

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	//fmt.Println(newTriangle(10, 20).area())

	//对上述三个图形的面积进行排序：首先构建shape类型的切片，后填充切片，最后实现排序
	shapes := make([]shape, 0, 3)
	shapes = append(shapes, c)
	shapes = append(shapes, t)
	shapes = append(shapes, r)

	fmt.Println("排序前: ", shapes[0].area(), shapes[1].area(), shapes[2].area())

	//升序
	sort.Slice(shapes, func(i, j int) bool {
		return shapes[i].area() > shapes[j].area()
	})

	fmt.Println("排序后: ", shapes[0].area(), shapes[1].area(), shapes[2].area())
}
