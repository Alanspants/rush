package main

import (
	"fmt"
	"math"
	"sort"
)

//定义面积接口
type areaer interface {
	GetArea() float32
}

//定义圆形结构体
type circle struct {
	radius float32
}

//定义三角形结构体
type triangle struct {
	name   string
	width  float32
	length float32
}

//定义长方形结构体
type rectangle struct {
	name   string
	width  float32
	length float32
}

type shape struct {
	name string
	s    float32
}

//圆形求面积方法
func (c *circle) GetArea() float32 {
	// return Pi * c.radius * c.radius
	return float32(math.Pi * math.Pow(float64(c.radius), 2))
}

//三角形求面积方法
func (t *triangle) GetArea() float32 {
	return 0.5 * (t.length * t.width)
}

//长方形求面积方法
func (r *rectangle) GetArea() float32 {
	return r.length * r.width
}

//求面积的函数
func area(a areaer) float32 {
	return a.GetArea()
}

//定义图形类型并实现 对应的sort  Interface接口
type shapeSlice []shape

func (a shapeSlice) Len() int           { return len(a) }
func (a shapeSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a shapeSlice) Less(i, j int) bool { return a[i].s < a[j].s }

func GetArea() {
	//实例化圆形 ，求面积
	c := &circle{radius: 4.0}
	fmt.Printf("圆的半径为: %f , 面积为: %0.2f\n", c.radius, area(c))

	//实例化三角形 ，求面积
	t := &triangle{width: 4.0, length: 4.0}
	fmt.Printf("三角形长为：%f , 宽为：%f ，面积为:  %0.2f\n", t.length, t.width, area(t))
	//实例化长方形 ， 求面积
	r := &rectangle{width: 4.0, length: 4.0}
	fmt.Printf("长方形长为：%f, 宽为：%f, 面积为: %0.2f\n", r.length, r.width, area(r))

}

func ShapeSort() {
	//实例化圆形
	c := &circle{radius: 4.0}
	//实例化三角形
	t := &triangle{width: 4.0, length: 4.0}
	//实例化长方形
	r := &rectangle{width: 4.0, length: 4.0}

	//升序排序
	fmt.Println("~~~~~~~~~~升序排序~~~~~~~~~~~~~")
	shapes1 := make([]shape, 0, 3)
	shapes1 = append(shapes1, shape{name: "圆形", s: area(c)}, shape{name: "三角形", s: area(t)}, shape{name: "长方形", s: area(r)})
	fmt.Println("升序排序前：", shapes1)
	sort.Slice(shapes1, func(i, j int) bool {
		return shapes1[i].s < shapes1[j].s
	})
	fmt.Println("升序排序后", shapes1)
	//降序排列
	fmt.Println("~~~~~~~~~~~降序排列~~~~~~~~~~~~")
	shapes := shapeSlice{}
	shapes = append(shapes, shape{name: "圆形", s: area(c)}, shape{name: "三角形", s: area(t)}, shape{name: "长方形", s: area(r)})
	fmt.Println("降序排序前：", shapes)
	sort.Sort(sort.Reverse(shapes))
	fmt.Println("降序排序后：", shapes)

}

func main() {
	// 作业
	// 1、实现对圆形、circle 三角形 triangle、长方形求面积rectangle。
	// 2、利用第1题，构造3个以上图形，至少圆形、三角形、矩形各有一个，对上题的图形按照面积降序排列

	//求所有图形面积
	GetArea()
	//排序
	ShapeSort()

}
