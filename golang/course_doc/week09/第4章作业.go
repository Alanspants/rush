package main

import (
	"fmt"
	"math"
	"sort"
)

type Interface interface {
	Area() float32
}

type Shape struct {
	area float32 // 面积属性，记录计算过的面积结果
}

func (s *Shape) String() string {
	return fmt.Sprintf("%T = %v", s, s.area)
}

type Circle struct {
	Shape
	r float32
}

func NewCircle(r float32) *Circle {
	return &Circle{r: r}
}

func (c *Circle) Area() float32 {
	if c.area == 0 {
		c.area = math.Pi * c.r * c.r
	}
	return c.area
}

type Triangle struct {
	Shape
	base, height float32
}

func NewTriangle(base, height float32) *Triangle {
	return &Triangle{base: base, height: height}
}

func (t *Triangle) Area() float32 {
	if t.area == 0 {
		t.area = t.base * t.height / 2
	}
	return t.area
}

type Rectangle struct {
	Shape
	width, height float32
}

func NewRectangle(width, height float32) *Rectangle {
	return &Rectangle{Shape{}, width, height}
}

func (r *Rectangle) Area() float32 {
	if r.area == 0 {
		r.area = r.width * r.height
	}
	return r.area
}

// func (r *Rectangle) String() string {
// 	return fmt.Sprintf("%T %.2f", r, r.Area())
// }

func main() {
	c := NewCircle(4)
	fmt.Println(c, c.Area())
	t := NewTriangle(2, 3)
	fmt.Println(t, t.Area())
	r := NewRectangle(2, 3)
	fmt.Println(r, r.Area())
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// 从接口角度出发，可以看到他们有共同的特征
	var i Interface = c
	fmt.Println(i.Area())
	i = t
	fmt.Println(i.Area())

	var shapes = []Interface{c, t, r}
	fmt.Println(shapes)
	sort.Slice(shapes, func(i, j int) bool {
		return shapes[i].Area() > shapes[j].Area()
	}) // Slice Sort 内部都要调用排序函数，内部有排序算法
	// 高阶函数，排序算法不知道你有多长、如何swap、如何比较大小
	// 提供给排序函数3个函数，需要3个函数类型形参
	// Go不需要你提供3个函数形参，改成了接口，符合该接口要求就实现了3个接口函数
	// 排序函数内部就可以直接对该接口类型数据直接调用这3个方法
	fmt.Println(shapes) // Stringer

	for i, v := range shapes {
		fmt.Println(i, v, v.Area())
	}

}
