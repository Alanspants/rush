package main

import (
	"fmt"
	"math"
	"sort"
)

// ExceptionHandle 异常处理
func ExceptionHandle() {
	fmt.Println("\n··········································开始捕获异常···········································")
	err := recover()
	if err != nil {
		fmt.Printf("\n~~~~%T~~~~%#[1]v``````````````````````\n", err)
	} else {
		fmt.Println("@@@@@@@@@@@有惊无险！##########")
	}
	fmt.Println("\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!异常处理结束!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()
}

// AcreageCalculator 通用接口
type AcreageCalculator interface {
	Acreage() float64
}

// Calculate 通用计算面积公式 多态
func Calculate(shape AcreageCalculator) {
	shape.Acreage()
}

type Shape struct {
	Name      string
	Perimeter float64 //周长
	arrea     float64 //面积
}

func (s *Shape) String() string {
	return fmt.Sprintf("%s的面积:%.2f ", s.Name, s.arrea)
}

// NewShape 构造方法
func NewShape(name string, perimeter, acreage float64) *Shape {
	return &Shape{Name: name, Perimeter: perimeter, arrea: acreage}
}

// Round 圆形结构体
type Round struct {
	Shape            // 半径 radius==Perimeter/2π
	diameter float64 //直径
}

// NewRound 圆形构造方法
func NewRound(s Shape, d float64) *Round {
	return &Round{s, d}
}

// Acreage 对圆求面积 Round's Acreage=πr²
func (r *Round) Acreage() float64 {
	if r.arrea == 0 { //判断一下，避免重复计算
		r.arrea = math.Pi * math.Pow(r.diameter/2, 2) // Acreage=(perimeter/2)*radius
	}
	return r.arrea
}

// Triangle 三角形结构体
type Triangle struct {
	Shape
	a, b, c float64 //三条边
}

// NewTriangle 构造三角形
func NewTriangle(s Shape, a, b, c float64) *Triangle {
	return &Triangle{s, a, b, c}
}

// Acreage 对三角形求面积 Triangle's S=√p(p-a)(p-b)(p-c) 海伦公式
func (t *Triangle) Acreage() float64 {
	if t.a+t.b <= t.c || t.a+t.c <= t.b || t.b+t.c <= t.a {
		defer ExceptionHandle()
		panic("三角形两条边之和必须大于第三边，否则无法构成三角形！")
	}
	if t.a+t.b+t.c != t.Perimeter {
		defer ExceptionHandle()
		panic("传参有误，三条边之和不等于周长,计算面积为0！")
	}
	if t.arrea == 0 { //避免重复计算
		p := t.Perimeter / 2
		r := p * (p - t.a) * (p - t.b) * (p - t.c)
		t.arrea = math.Sqrt(r)
	}
	return t.arrea
}

// Square 正方形结构体
type Square struct {
	Shape
	edge float64
}

// NewSquare 正方形构造方法
func NewSquare(s Shape, e float64) *Square {
	return &Square{s, e}
}

// Acreage 对正方形求面积 Square's Acreage=side*side
func (s *Square) Acreage() float64 {
	if s.Perimeter/4 != s.edge {
		defer ExceptionHandle()
		panic("正方形周长必须是每条边的四倍！")
	}
	if s.arrea == 0 { //重复计算判断
		s.arrea = math.Pow(s.edge, 2)
	}
	return s.arrea
}

// Rectangle 矩形结构体
type Rectangle struct {
	Shape
	floor float64
}

// Acreage 对矩形求面积 Rectangle's Acreage=(perimeter/2-floor)*floor
func (t *Rectangle) Acreage() float64 {
	if t.floor >= t.Perimeter/2 {
		defer ExceptionHandle()
		panic("矩形边长必须小于周长的一半！")
	}
	// 判断有无计算过，因为初始值是0
	if t.arrea == 0 {
		t.arrea = t.floor * (t.Perimeter/2 - t.floor)
	}
	return t.arrea
}

// ShapeSlice 实现了sort.Sort.Interface接口，用做强制类型转换;[]Shape切片可用来存放circle、triangle、square、rectangle结构体
type ShapeSlice []Shape

func (a ShapeSlice) Len() int           { return len(a) }
func (a ShapeSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ShapeSlice) Less(i, j int) bool { return a[i].arrea < a[j].arrea } //默认升序

func main() {
	// 1、实现对圆形、三角形、长方形求面积。
	// 圆形实例
	circle := NewRound(*NewShape("〇圆形", 36, 0), 36/math.Pi)
	Calculate(circle) //计算circle面积
	fmt.Printf("%T %+[1]v\n", circle)
	fmt.Println("\n——————————————————————————————————————————下面打印三角形对象————————————————————————————————————————————————————————")
	// 三角形实例化
	triangle := NewTriangle(*NewShape("①正三角形", 36, 0), 12, 12, 12)
	Calculate(triangle) //算正三角形面积
	fmt.Printf("%T %+[1]v\n", triangle)
	t2 := NewTriangle(*NewShape("②锐角三角形", 36, 0), 11, 12, 13)
	Calculate(t2) //算锐角三角形面积
	fmt.Printf("%T %+[1]v\n", t2)
	/* 	t3 := NewTriangle(*NewShape("③面积是0,无法构成三角形", 36, 0), 12, 13, 14)
	   	Calculate(t3) //无面积
	   	fmt.Printf("%T %+[1]v\n", t3) */
	fmt.Println("\n------------------------------------------输出正方形实例---------------------------------------------")
	//实例化正方形
	square := NewSquare(*NewShape("④正方形", 36, 0), 9)
	Calculate(square) //算面积
	fmt.Printf("%T %+[1]v\n", square)
	fmt.Println("\n===========================================下面打印矩形对象==========================================================")
	//矩形实例化
	rectangle := &Rectangle{*NewShape("⑥长矩形", 36, 0), 17}
	Calculate(rectangle) //计算长方形面积
	fmt.Printf("%T %+[1]v\n", rectangle)
	r2 := &Rectangle{*NewShape("⑤宽矩形", 36, 0), 10}
	Calculate(r2) //计算面积
	fmt.Printf("%T %+[1]v\n", r2)
	// 2、利用第1题，构造3个以上图形，至少圆形、三角形、矩形各有一个，对上题的图形按照面积降序排列
	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~Sorting start对相同周长的形状按面积ASCending~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	//用结构体切片[]Shape升序
	ss := make([]Shape, 0, 6)
	// fmt.Printf("%T==> %+[1]v\t %[2]T==> %+[2]v\n", ss, ShapeSlice(ss))
	ss = append(ss, circle.Shape, triangle.Shape, t2.Shape, square.Shape, rectangle.Shape, r2.Shape)
	//升序 传入实现了sort包里的接口Interface 的 ShapeSlice
	sort.Sort(ShapeSlice(ss)) //sort.Sort(data sort.Interface)参数为sort包下的Interface类型，Interface定义了三个方法：type Interface interface { Len() int  Less(i, j int) bool  Swap(i, j int) };ShapeSlice实现了Len() int  Less(i, j int) bool  Swap(i, j int)这三个方法，等于实现了接口Interface
	fmt.Println(ss)
	fmt.Println("\nDESC+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++DESCending")
	//降序
	var shapes = []AcreageCalculator{circle, triangle, t2, square, rectangle, r2}
	sort.Slice(shapes, func(i, j int) bool {
		return shapes[i].Acreage() > shapes[j].Acreage()
	})
	fmt.Printf("%+v\n", shapes)
}
