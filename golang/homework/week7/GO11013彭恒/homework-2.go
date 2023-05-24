package main

import (
	"fmt"
	"math"
	"sort"
)

// PI 全局常量PI，用于求圆形面积
const PI = 3.14

// shape 定义形状接口，方法area求面积，...float64是float64类型的可变形参
type shape interface {
	area(args ...float64) float64
}

type entry struct {
	key   string
	value float64
}

// 定义圆形结构体，属性radius是半径
type round struct {
	radius float64
}

// 定义三角形结构体，属性bottom是底，属性high是高
type triangle struct {
	bottom, high float64
}

// 定义长方形结构体，属性length是长，属性width是宽
type rectangle struct {
	length, width float64
}

// 结构体round实现了shape接口，重写了area方法
func (*round) area(radius float64) float64 {
	return PI * math.Pow(radius, 2)
}

// 结构体triangle实现了shape接口，重写了area方法
func (*triangle) area(bottom, high float64) float64 {
	return bottom * high / 2
}

// 结构体rectangle实现了shape接口，重写了area方法
func (*rectangle) area(length, width float64) float64 {
	return length * width
}

func main() {
	// 实例化圆形，调用圆形area方法求面积
	roundArea := new(round).area(4.6)
	fmt.Printf("round     area is: %f\n", roundArea)
	// 实例化三角形，调用三角形area方法求面积
	triangleArea := new(triangle).area(3.56, 5.93)
	fmt.Printf("triangle  area is: %f\n", triangleArea)
	// 实例化长方形，调用长方形area方法求面积
	rectangleArea := new(rectangle).area(3.3, 4.7)
	fmt.Printf("rectangle area is: %f\n", rectangleArea)

	// 定义map m存放圆形、三角形、长方形面积
	m := make(map[string]float64)
	m["roundArea"] = roundArea
	m["triangleArea"] = triangleArea
	m["rectangleArea"] = rectangleArea

	// 定义长度、容量都为map m元素entry切片，存放各个图形面积
	entries := make([]entry, len(m))
	i := 0
	// 遍历map m，k为图形，v为面积
	for k, v := range m {
		// 存放键为i，值为entry实例到entries切片中
		entries[i] = entry{k, v}
		i++
	}
	// 使用sort.Slice方法对切片进行排序，> 为降序排列
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].value > entries[j].value
	})
	// 输出排序后的entry切片
	fmt.Printf("Sort by area in descending order: %v", entries)
}
