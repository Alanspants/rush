package main

import (
	"fmt"
	"sort"
)

//定义Area接口 用于多态实现
type Areaer interface {
	area() int
}

//圆形
type Yuan struct {
	pi float32
	r  int
}

//长方形
type Chang struct {
	w, l int
}

//三角形
type San struct {
	Chang //匿名嵌套长方形，实现继承的效果
}

//第二题构建三个以上图形排序的结构体
type AllGraph struct {
	name string
	area int
}

//第二题使用，构造函数
func NewAllGraph(name string, area int) *AllGraph {
	return &AllGraph{name, area}
}

//第二题使用排序
type AllgraphSlice []AllGraph

func (x AllgraphSlice) Len() int      { return len(x) }
func (x AllgraphSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x AllgraphSlice) Less(i, j int) bool {
	//降序排序
	return x[i].area > x[j].area

}

//实现多态
func foo(a Areaer) int {
	//a.area()
	return a.area()
}

//求三种类型的面积*(圆 长方形 三角形)
func (y *Yuan) area() int {
	area := y.pi * float32(y.r*y.r)
	//fmt.Printf("圆形面积是：%.f\n", area)
	return int(area)
}
func (c *Chang) area() int {
	area := c.w * c.l
	//fmt.Printf("长方形面积是：%d\n", area)
	return area
}
func (s *San) area() int {
	area := s.w * s.l / 2
	//fmt.Printf("三角形面积是：%d\n", area)
	return area
}

func main() {
	//第一题求面积
	//圆形面积 pi*r^2
	y1 := &Yuan{3.14, 3}
	//y1.area()
	fmt.Printf("圆形面积是：%d\n", y1.area())

	//长方形面积 长*宽
	c1 := new(Chang)
	c1.w = 3
	c1.l = 4
	//c1.area()
	fmt.Printf("长方形面积是：%d\n", c1.area())

	//三角形面积 底*高/2 ,三角形继承长方形,多态实现求面积
	s1 := &San{}
	s1.w = 10
	s1.Chang.l = 4
	//foo(s1)
	fmt.Printf("三角形面积是：%d\n", foo(s1))

	//第二题对面积排序
	//构建三个实例
	n1 := NewAllGraph("Yuan", foo(y1))
	n2 := NewAllGraph("Chang", foo(c1))
	n3 := NewAllGraph("San", foo(s1))
	alls := []AllGraph{*n1, *n2, *n3}
	fmt.Println("排序前：", alls)
	//对结构体 降序排序
	sort.Sort(AllgraphSlice(alls))
	fmt.Println("排序后：", alls)

}
