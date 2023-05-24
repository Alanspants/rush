package main

import (
	"fmt"
	"math"
)

// PI 全局常量PI，用于求圆形面积
const PI = 3.14

// shape 定义形状接口，方法area求面积，...float64是float64类型的可变形参
type shape interface {
	area(args ...float64) float64
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
}
