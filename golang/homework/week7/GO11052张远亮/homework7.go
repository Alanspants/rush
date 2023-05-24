package main

import (
	"fmt"
	"sort"
)

/*
 1、实现对圆形、三角形、长方形求面积。
*/

// 圆形
type circular struct {
	//半径
	radius float64
}

var pai = 3.14

// 三角形
type triangle struct {
	//底
	bottom float64
	//高
	tall float64
}

// 长方形
type rectangle struct {
	//长
	long float64
	//宽
	wide float64
}

// 定义面积的接口
type newArea interface {
	area() float64
}

// 求圆面积的方法
func (c *circular) area() float64 {
	//pai*半径的平方
	area := pai * c.radius * c.radius
	return area
}

// 求三角形面积的方法
func (t *triangle) area() float64 {
	//底*高/2
	area := (t.bottom * t.tall) / 2
	return area
}

// 求长方形面积的方法
func (r *rectangle) area() float64 {
	//长*高
	area := r.long * r.wide
	return area
}

// 圆形的构造函数
func newCircular(radius float64) *circular {
	return &circular{radius: radius}
}

// 三角形的构造函数
func newTriangle(bottom, tall float64) *triangle {
	return &triangle{
		bottom: bottom,
		tall:   tall,
	}
}

// 长方形的构造函数
func newRectangle(long, wide float64) *rectangle {
	return &rectangle{
		long: long,
		wide: wide,
	}
}

//实现String接口
func (c *circular) String() string {
	return fmt.Sprintf("%.2f",c.area())
}

func (t *triangle) String() string {
	return fmt.Sprintf("%.2f",t.area())
}
func (r *rectangle) String() string {
	return fmt.Sprintf("%.2f",r.area())
}

func main() {
	//创建存放所有形状面积的切片
	areas := make([]float64, 0, 3)
	//求圆形面积
	c1 := newCircular(5)
	ret1 := c1.area()
	areas = append(areas, ret1)
	fmt.Println(ret1)
	//求三角形面积
	t1 := newTriangle(10, 5)
	ret2 := t1.area()
	fmt.Println(ret2)
	areas = append(areas, ret2)
	//求长方形面积
	r1 := newRectangle(10, 5)
	ret3 := r1.area()
	fmt.Println(ret3)
	areas = append(areas, ret3)
	//面积的升序排序
	sort.Float64s(areas)
	fmt.Printf("升序：%v\n", areas)
	//面积的降序排序
	sort.Sort(sort.Reverse(sort.Float64Slice(areas)))
	fmt.Printf("降序:%v\n", areas)
	//面积的升序排序2  从接口角度出发，可以看到他们有共同的特征
	//var i newArea = c1
	//fmt.Println(i.area)
	//i = t1
	//fmt.Println(i.area())
	var n newArea
	n = &circular{radius:10}
	fmt.Println(n)

	var shapes  = []newArea{c1,t1,r1}
	sort.Slice(shapes, func(i, j int) bool {
		return shapes[i].area() < shapes[j].area()
	})
	for i, v := range shapes {
		fmt.Println(i,v.area())
	}

	fmt.Println(shapes)

}
