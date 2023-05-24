package main

import (
	"fmt"
	"math"
	"sort"
)

//实现对圆形 三角形 长方形 求面积
//2.构造3个以上图形，至少每个各一个对上题图形按照面积降序排列

type Circle struct {
	radius float64 //圆形半径
	area   float64
}

// NewCircle 构造函数 输入半径得到一个圆的结构体实例
func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

// GetArea 返回圆形的面积
func (c *Circle) GetArea() (area float64) {
	c.area = math.Pi * math.Pow(c.radius, 2) //pow是次方 sqpt 是开方
	return
}

type Triangle struct {
	edge1 float64
	edge2 float64
	edge3 float64
	area  float64
}

// NewTriangle 构造函数创建一个三角形实例
func NewTriangle(edge1 float64, edge2 float64, edge3 float64) *Triangle {
	return &Triangle{edge1: edge1, edge2: edge2, edge3: edge3}
}

// GetArea 返回三角形面积
func (t *Triangle) GetArea() (area float64) {
	//S=sqrt[p(p-a)(p-b)(p-c)]//海伦公式
	p := t.edge1 + t.edge2 + t.edge3 //p表示三角形的周长
	t.area = math.Sqrt(p * (p - t.edge1) * (p - t.edge2) * (p - t.edge3))
	return
}

type Rectangle struct {
	length float64
	wide   float64
	area   float64
}

// NewRectangle 构造函数返回长方形的实例
func NewRectangle(length float64, wide float64) *Rectangle {
	return &Rectangle{length: length, wide: wide}
}

// GetArea 返回长方形实例的面积
func (r *Rectangle) GetArea() (area float64) {
	r.area = r.wide * r.length
	return
}

// Area  用于排序
type Area struct {
	name string
	area float64
}

//AreaSlice 面积类型切片 用来实现sort.sort方法
type AreaSlice []Area

func (x AreaSlice) Len() int           { return len(x) }
func (x AreaSlice) Less(i, j int) bool { return x[i].area < x[j].area }
func (x AreaSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	//创建map序列用来保存图形的面积
	MapArea := make(map[string]float64, 3)

	circle1 := NewCircle(3) //创建一个圆实例
	circle1.GetArea()
	MapArea["circle1"] = circle1.area
	fmt.Printf("%#v\n", circle1)

	triangle1 := NewTriangle(1, 2, 3)
	triangle1.GetArea()
	fmt.Printf("%#v\n", triangle1)
	MapArea["triangle1"] = triangle1.area

	rectangle1 := NewRectangle(2, 20)
	rectangle1.GetArea()
	MapArea["rectangle1"] = rectangle1.area
	fmt.Printf("%#v\n", rectangle1)
	fmt.Printf("%#v\n", MapArea)
	//接下来对结构体的value进行排序
	//将map的kv,放到一个[]Area类型中

	//创建一个切片
	AreaSlice1 := make([]Area, 0, len(MapArea))

	//将map的kv放到切片
	for k, v := range MapArea {
		AreaSlice1 = append(AreaSlice1, Area{k, v})
	}

	fmt.Println("排序前")
	fmt.Printf("%v\n", AreaSlice1)
	sort.Sort(sort.Reverse(AreaSlice(AreaSlice1)))
	fmt.Println("按照面积降序排序后")
	fmt.Printf("%v\n", AreaSlice1)
}
