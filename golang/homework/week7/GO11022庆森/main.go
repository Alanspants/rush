package main

import (
	"fmt"
	"math"
	"sort"
)

type circle struct {
	r float64
}

type triangle struct {
	bottom float64
	high   float64
}

type rectangle struct {
	length float64
	width  float64
}

func (a circle) area() float64 {
	return a.r * a.r * math.Pi
}

func (b triangle) area() float64 {
	return b.bottom * b.high / 2
}

func (c rectangle) area() float64 {
	return c.length * c.width
}

type areaer interface {
	area() float64
}

func main() {
	a := circle{r: 5}
	b := triangle{5, 5}
	c := rectangle{5, 6}
	s1 := a.area()
	s2 := b.area()
	s3 := c.area()
	fmt.Println(s1, s2, s3)

	areaslice := make([]areaer, 0, 3)
	areaslice = append(areaslice, a)
	areaslice = append(areaslice, b)
	areaslice = append(areaslice, c)
	fmt.Println("排序前", areaslice)

	sort.Slice(areaslice, func(i, j int) bool {
		return areaslice[i].area() > areaslice[j].area()
	})
	fmt.Println("排序后", areaslice)
}
