package main

import (
	"fmt"
	"math"
	"sort"
)

type triangle struct {
	b float64 //底边长
	h float64 //高
}

type circle struct {
	r float64 //半径
}

type rectangle struct {
	x, y float64 //长、宽
}

func (t triangle) area() float64 {
	return t.b * t.h * 0.5
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r rectangle) area() float64 {
	return r.x * r.y
}

type areaer interface {
	area() float64
}

type graphicArea struct {
	areaer
	area float64
}

type areaSlice []graphicArea

func (a areaSlice) Len() int           { return len(a) }
func (a areaSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a areaSlice) Less(i, j int) bool { return a[i].area < a[j].area }

func main() {

	t1 := triangle{10, 5}
	t2 := triangle{6, 5}
	c1 := circle{6}
	c2 := circle{5}
	r1 := rectangle{7, 8}
	r2 := rectangle{3, 4}

	var s1 = []areaer{t1, t2, c1, c2, r1, r2}
	s2 := make(areaSlice, len(s1))
	for i, v := range s1 {
		s2[i].areaer = v
		s2[i].area = v.area()
	}

	fmt.Println("排序前：", s2)
	sort.Sort(sort.Reverse(s2))
	fmt.Println("排序后：", s2)

}
