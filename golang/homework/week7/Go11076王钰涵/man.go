package main

import (
	"fmt"
	"sort"
)

//圆形，Πr2或者Πd
//三角形，底×高除2
//长方形，长×宽

//求面积接口
type Squaring interface {
	square() float32
}

//所有形状的父类（面向对象这样叫），图形
type tx struct {
	txName string
}

//圆形
type circle struct {
	tx
	radius float32
}

//长方形
type rectangle struct {
	tx
	length float32
	width  float32
}

//三角形
type triangle struct {
	tx
	Base   float32
	height float32
}

//面积的结构体
type mjS struct {
	Name   string
	square float32
}

//圆形构造函数
func YConstruction(txName string, square float32) *circle {
	return &circle{tx{txName}, square}
}

//圆形实现接口
func (c *circle) square() float32 {
	return 3.14 * c.radius * c.radius
}

//三角形构造函数
func SConstruction(txName string, Base float32, height float32) *triangle {
	//return &triangle{Name, Base, height}
	return &triangle{tx{txName}, Base, height}
}

//三角形实现接口
func (c *triangle) square() float32 {
	return c.Base * c.height / 2
}

//长方形构造函数
func CjConstruction(txName string, length float32, width float32) *rectangle {
	return &rectangle{tx{txName}, length, width}
}

//长方形实现接口
func (r *rectangle) square() float32 {
	return r.length * r.width
}

func mj(s Squaring) float32 {
	m := s.square()
	switch s.(type) {
	case *circle:
		fmt.Printf("圆的面积为：%f\n", m)
	case *rectangle:
		fmt.Printf("长方形的面积为：%f\n", m)
	case *triangle:
		fmt.Printf("三角形的面积为：%f\n", m)
	}
	return m
}

func main() {
	//定义个map存放类型和面积
	var Mmap = make(map[string]float32)
	//定义个slice存放map生成数据
	var mSlace = make([]mjS, 0, 20)
	//定义圆形实例
	y1 := YConstruction("y1", 2)
	y2 := YConstruction("y2", 4)
	//定义三角形实例
	s1 := SConstruction("s1", 2, 3)
	s2 := SConstruction("s2", 3, 7)
	//定义长方形实例
	c1 := CjConstruction("c1", 4, 5)
	c2 := CjConstruction("c2", 4, 8)
	//根据面积生成map
	Mmap[y1.txName] = mj(y1)
	Mmap[y2.txName] = mj(s1)
	Mmap[c1.txName] = mj(c1)
	Mmap[c2.txName] = mj(y2)
	Mmap[s1.txName] = mj(s1)
	Mmap[s2.txName] = mj(s2)
	fmt.Printf("生成map存储面积： %v\n", Mmap)
	//定义面积的struct，并根据struct创建struct切片
	for k, v := range Mmap {
		m := mjS{Name: k, square: v}
		mSlace = append(mSlace, m)
	}
	fmt.Printf("将map生成slice，此为初始slice： %v\n", mSlace)
	//根据切片排序
	sort.Slice(mSlace, func(i, j int) bool {
		return mSlace[i].square > mSlace[j].square
	})
	fmt.Printf("排序后的slice： %v\n", mSlace)
}
