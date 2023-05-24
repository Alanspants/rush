package main

import (
	"fmt"
	"math"
	"sort"
)

// 计算面积接口
type CalcArea interface {
	Area() float64
}

// 三角形
type Triangle struct {
	lenth float64
	width float64
}

// 长乘宽除以2
func (t Triangle) Area() float64 {
	return float64(t.lenth * t.width / 2)
}

func (t Triangle) String() string {
	return fmt.Sprintf("%f", t.Area())
}

func NewTriangle(l, w float64) *Triangle {
	return &Triangle{l, w}
}

type Square struct {
	lenth float64
	width float64
}

// 长乘宽
func (s Square) Area() float64 {
	return s.lenth * s.width
}

func (s Square) String() string {
	return fmt.Sprintf("%f", s.Area())
}

func NewSquare(l, w float64) *Square {
	return &Square{l, w}
}

type Round struct {
	diameter float64
	pi       float64
}

// 如果Pi未指定使用math.pi
func (r Round) Area() float64 {
	if r.pi == 0 {
		r.pi = math.Pi
	}
	return r.pi * math.Pow(r.diameter, 2)
}

func (r Round) String() string {
	return fmt.Sprintf("%f", r.Area())
}

func NewRound(d, pi float64) *Round {
	return &Round{d, pi}
}

type CalcAreaSlice []CalcArea

func (c CalcAreaSlice) Len() int           { return len(c) }
func (c CalcAreaSlice) Less(i, j int) bool { return c[i].Area() > c[j].Area() }
func (c CalcAreaSlice) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func main() {
	// 习题一
	//	圆形
	var r CalcArea = NewRound(10, 0)
	// 圆形面积
	fmt.Println(r.Area())

	//	方形
	var s CalcArea = NewSquare(10, 10)
	//	方形面积
	fmt.Println(s.Area())

	//	三角形
	var t CalcArea = NewTriangle(10, 10)
	//	三角形面积
	fmt.Println(t.Area())

	//	习题二
	//	对r,s,t降序排列
	var areaSlices = []CalcArea{r, s, t}
	var s1 = CalcAreaSlice(areaSlices)

	sort.Sort(s1)
	fmt.Printf("%+v %[1]T\n", areaSlices)
	fmt.Printf("%+v %[1]T\n", s1)
}
