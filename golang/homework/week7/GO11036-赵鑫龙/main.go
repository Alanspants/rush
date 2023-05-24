package main

import (
	"fmt"
	"math"
	"sort"
)

// Areaer 实现一个Area接口
type Areaer interface {
	area() float64
}

// 为三角形，圆形，长方形定义一个结构体
type Triangle struct {
	a, b, c float64
}

type Round struct {
	radius float64
	pi     float64
}

type Rectangle struct {
	long  float64
	width float64
}

// 为三角形，圆形长方形定义一个area方法
func (r *Round) area() float64 {
	s := r.pi * math.Pow(r.radius, 2)
	//fmt.Printf("圆形半径为%v\n", r.radius)
	return s
}
func (t *Triangle) area() float64 {
	s := (t.a + t.b + t.c) / 2
	a := math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
	//fmt.Printf("三角形三边长为%v,%v,%v\n", t.a, t.b, t.c)
	//fmt.Printf("(%v + %v + %v)/2 = %v\n", t.a, t.b, t.c, s)
	//fmt.Printf("三角形的面积为√(%v * (%[1]v-%[2]v)*(%[1]v-%[3]v)*(%[1]v-%[4]v)) = %.2f\n", s, t.a, t.b, t.c, a)
	return a
}
func (r *Rectangle) area() float64 {
	s := r.long * r.width
	//fmt.Printf("长方形长为%v,宽为%v\n", r.long, r.width)
	//fmt.Printf("长方形的面积为: %v * %v = %v\n", r.long, r.width, s)
	return s
}

// TriangleFunc 为三角形结构体定义一个函数
func TriangleFunc(a, b, c float64) *Triangle {
	x := &Triangle{a, b, c}
	return x
}

// RoundFunc 为圆形结构体定义一个函数
func RoundFunc(pi float64, radius float64) *Round {
	x := &Round{radius, pi}
	return x
}

// RectangleFunc 为长方形结构体定义一个函数
func RectangleFunc(long, width float64) *Rectangle {
	x := &Rectangle{long, width}
	return x
}

// Area 定义一个结构函数，调用Area接口
//func Area(a Areaer) float64 {
//	return a.area()
//}

func main() {
	shapes := []Areaer{
		TriangleFunc(5, 5, 4),
		RoundFunc(3.14, 6),
		RectangleFunc(5, 6),
		RectangleFunc(10, 11),
		RectangleFunc(20, 6),
	}

	sort.Slice(shapes, func(i, j int) bool {
		return shapes[i].area() > shapes[j].area()
	})

	fmt.Println("从高到低排序:")
	for _, shape := range shapes {
		switch s := shape.(type) {
		case *Triangle:
			fmt.Printf("三角形: %.2f\n", s.area())
		case *Round:
			fmt.Printf("圆形: %.2f\n", s.area())
		case *Rectangle:
			fmt.Printf("长方形: %.2f\n", s.area())
		}
	}
}
