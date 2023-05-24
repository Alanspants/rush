package main

import (
	"fmt"
	"sort"
)

// 1、实现对圆形、三角形、长方形求面积。
// 2、利用第1题，构造3个以上图形，至少圆形、三角形、矩形各有一个，对上题的图形按照面积降序排列

//圆的面积=圆周率×半径×半径，S=πr²
// 1.三角形面积=1/2×底×高;或者说，三角形面积=(底×高)÷2;
// 2.已知三角形三边a,b,c，则(海伦公式)(S=(a+b+c)/2)。
//长方形的面积=长×宽，S=ab。

type Area struct {
	s float64
}

//圆形
type round struct {
	Area
	p, r float64
}

//三角形
type triangle struct {
	Area
	d, h float64
}

//长方型
type rectangle struct {
	Area
	a, b float64
}

//圆形构造函数
func newround(p, r float64) round {
	s := p * r * r
	rr := round{Area{s}, p, r}
	return rr
}

//三角形构造函数
func newtriangle(d, h float64) triangle {
	s := d * h / 2
	rr := triangle{Area{s}, d, h}
	return rr
}

//长方型构造函数
func newrectangle(a, b float64) rectangle {
	s := a * b
	rr := rectangle{Area{s}, a, b}
	return rr
}

// type Areas interface {
// 	round() float64
// 	triangle() float64
// 	rectangle() float64
// }

// func (m Area) round() float64 {
// 	return m.b * m.c * m.c
// }

// func (m Area) triangle() float64 {
// 	return m.b * m.c / 2
// }

// func (m Area) rectangle() float64 {
// 	return m.b * m.c
// }

func main() {
	//第一题
	//圆形
	fmt.Println(newround(3.14, 5))
	//三角形
	fmt.Println(newtriangle(4, 6))
	//长方形
	fmt.Println(newrectangle(6, 7))

	//第二题
	s1 := newround(3.14, 8)
	s2 := newtriangle(2, 4)
	s3 := newrectangle(4, 6)

	// sss := make([]round, 0)
	// fmt.Println(sss)
	// sss = append(sss, s1)
	// fmt.Println(sss)

	s := make(map[float64]map[float64]float64, 0)

	ss1 := make(map[float64]float64, 0)
	ss1[s1.p] = s1.r
	s[s1.s] = ss1

	ss2 := make(map[float64]float64, 0)
	ss2[s2.d] = s2.h
	s[s2.s] = ss2

	ss3 := make(map[float64]float64, 0)
	ss3[s3.a] = s3.b
	s[s3.s] = ss3

	var sli = make([]float64, 0)
	for k, _ := range s {
		sli = append(sli, k)
	}
	sort.Float64s(sli)
	for i := len(sli) - 1; i >= 0; i-- {
		fmt.Println(sli[i], s[sli[i]])
	}

}
