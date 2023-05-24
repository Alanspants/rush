package main

import (
	"fmt"
	"sort"
)

const Pi = 3

type calcArea interface {
	calc() int
}
type Circle struct {
	r int
}
type Triangle struct {
	bottom int
	height int
}
type Oblong struct {
	wide int
	long int
}

func (p *Circle) calc() int {
	return Pi * p.r * p.r
}

func (p *Triangle) calc() int {
	return p.bottom * p.height
}
func (p *Oblong) calc() int {
	return p.wide * p.long
}
func calc(c calcArea) int {
	return c.calc()
}

type aeraSort []calcArea

func (a aeraSort) Len() int           { return len(a) }
func (a aeraSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a aeraSort) Less(i, j int) bool { return a[i].calc() > a[j].calc() }
func main() {
	var a1 = &Circle{r: 3}
	var a2 = &Triangle{bottom: 3, height: 10}
	var a3 = &Oblong{wide: 1, long: 5}
	var a4 = &Oblong{wide: 2, long: 5}
	fmt.Println(calc(a1))
	fmt.Println(calc(a2))
	fmt.Println(calc(a3))
	fmt.Println(calc(a4))
	var s = aeraSort{a1, a2, a3, a4}
	sort.Sort(s)
	for _, v := range s {
		fmt.Printf("形状%T ----面积 %v\n", v, v.calc())
	}
}
