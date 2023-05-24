package main

import (
	"fmt"
	"math"
	"sort"
)

//1、实现对圆形、三角形、长方形求面积。
//2、利用第1题，构造3个以上图形，至少圆形、三角形、矩形各有一个，对上题的图形按照面积降序排列

type Interface interface {
	Area() float32
}

type Shap struct {
	area float32 //面积属性，记录计算过的面积结果，做为父类
}

func (r *Shap) String() string { //Stringer格式化输出
	return fmt.Sprintf("%.2f", r.area)
}

type Circle struct {
	Shap
	r float32
}

func (c *Circle) Area() float32 {
	if c.area == 0 {
		c.area = math.Pi * c.r * c.r
	}
	return c.area
}

// 构造函数
func NewCircle(r float32) *Circle {
	return &Circle{r: r}
}

type Triangle struct {
	Shap
	base, height float32
}

func (t *Triangle) Area() float32 {
	if t.area == 0 {
		t.area = t.base * t.height / 2
	}
	return t.area
}
func NewTriangle(base, height float32) *Triangle {
	return &Triangle{base: base, height: height}
}

type Rectangle struct {
	Shap
	width, height float32
}

func (r *Rectangle) Area() float32 {
	if r.area == 0 {
		r.area = r.width * r.height
	}
	return r.area
}
func NewRectangle(width, height float32) *Rectangle {
	return &Rectangle{width: width, height: height}
}

// func (r *Rectangle) String() string {
// 	return fmt.Sprintf("%2.f", r.area)
// }

func main() {
	c := NewCircle(4)
	fmt.Println(c, c.Area())
	t := NewTriangle(3, 4)
	fmt.Println(t, t.Area())
	r := NewRectangle(3, 4)
	fmt.Println(r, r.Area())
	fmt.Println("-----------------------")
	//从接口角度出发，可以看到他们有共同的特征
	var i Interface = c
	fmt.Println(i.Area())
	i = t
	fmt.Println(i.Area())

	var shapes = []Interface{c, t, r}
	fmt.Println(shapes[0], shapes[1], shapes[2]) //打印结果 没有排序
	sort.Slice(shapes, func(i, j int) bool {     //升序 排序
		return shapes[i].Area() < shapes[j].Area()
	}) //Slice Sort 内部都要调用排序函数，内部有排序算法
	//高阶函数，排序算法不知道你有多长，如何swap、如何比较大小
	//提供给排序函数3个函数，需要3个函数类型形参
	//Go不需要你提供3个函数形参，改成了接口，符合该接口要求就实现了3个接口函数
	//排序函数内部就可以直接对该接口类型数据直接调用这3个方法
	fmt.Println(shapes)        //Stringer 由于是指针地址在20行用Stringer格式化了输出
	for i, v := range shapes { //使用for循环打印每个值
		fmt.Println(i, v, v.Area())
	}
}
