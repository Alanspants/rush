package main

import (
	"fmt"
	"math"
	"sort"
)

type ShapeInterface interface {
	Area() float64
	GetAreaMethod() string
	GetName() string
}

func showArea(shape ShapeInterface) {
	fmt.Printf("%s  %s =  %f \n", shape.GetName(), shape.GetAreaMethod(), shape.Area())
}

type Shape struct {
	name string
	area float64
}

func (s *Shape) GetName() string {
	return s.name
}

type Rectangle struct {
	Shape
	w, h float64
}

func (r *Rectangle) Area() float64 {
	if r.area == -1 {
		r.area = r.w * r.h
	}
	return r.area
}
func (r *Rectangle) GetAreaMethod() string {
	return "w * h"
}

func NewRectangle(name string, w, h float64) *Rectangle {
	return &Rectangle{Shape: Shape{name: name, area: -1}, w: w, h: h}
}

type Circle struct {
	Shape
	r float64
}

func NewCircle(name string, r float64) *Circle {
	return &Circle{Shape: Shape{name: name, area: -1}, r: r}
}

func (c *Circle) Area() float64 {

	if c.area == -1 {
		c.area = c.r * c.r * math.Pi
	}
	return c.area
}

func (r *Circle) GetAreaMethod() string {
	return "r * r * math.Pi"
}

type Triangle struct {
	Shape
	a, h float64
}

func NewTriangle(name string, a, h float64) *Triangle {
	return &Triangle{Shape: Shape{name: name, area: -1}, a: a, h: h}
}

func (t *Triangle) Area() float64 {

	if t.area == -1 {
		t.area = t.a * t.h / 2
	}
	return t.area
}

func (r *Triangle) GetAreaMethod() string {
	return "a * h / 2"
}

type ShapeInterfaceSlice []ShapeInterface

func (x ShapeInterfaceSlice) Len() int           { return len(x) }
func (x ShapeInterfaceSlice) Less(i, j int) bool { return x[i].Area() < x[j].Area() }
func (x ShapeInterfaceSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {

	c := NewCircle("Circle", 10)
	r := NewRectangle("Rectangle", 5, 4)
	t := NewTriangle("Triangle", 5, 4)
	showArea(c)
	showArea(r)
	showArea(t)
	s1 := make([]ShapeInterface, 0, 3)
	s1 = append(s1, t)
	s1 = append(s1, c)
	s1 = append(s1, r)
	showS1(s1)
	// 升序
	sort.SliceStable(s1, func(i, j int) bool {
		return s1[i].Area() < s1[j].Area()
	})
	showS1(s1)
	// 降序
	sort.SliceStable(s1, func(i, j int) bool {
		return s1[i].Area() > s1[j].Area()
	})
	showS1(s1)
	// 升序
	sort.Sort(ShapeInterfaceSlice(s1))
	showS1(s1)

	// 降序
	sort.Sort(sort.Reverse(ShapeInterfaceSlice(s1)))
	showS1(s1)
}

func showS1(s1 []ShapeInterface) {
	for _, v := range s1 {
		fmt.Printf("%+v", v)
	}
	fmt.Println("--")
}
