package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type circular struct {
	radius float32
}
type triangle struct {
	low  float32
	high float32
}
type rectangle struct {
	wide float32
	high float32
}
type square interface {
	area() float32
}

func (c circular) area() float32 {
	a := math.Pi * c.radius * c.radius
	return a
}
func (t triangle) area() float32 {
	a := 0.5 * t.high * t.low
	return a
}
func (r rectangle) area() float32 {
	a := r.wide * r.high
	return a
}

func main() {
	var c square = &circular{10}
	c.area()
	var t square = &triangle{5, 4}
	t.area()
	var r square = &rectangle{5, 4}
	r.area()

	squares := make([]square, 0)
	s := rand.New(rand.NewSource(time.Now().UnixMicro()))
	for i := 0; i < 3; i++ {
		squares = append(squares, circular{s.Float32() * 10})
	}
	for i := 0; i < 3; i++ {
		squares = append(squares, triangle{s.Float32() * 10, s.Float32() * 10})
	}
	for i := 0; i < 3; i++ {
		squares = append(squares, triangle{s.Float32() * 10, s.Float32() * 10})
	}
	for _, i := range squares {
		fmt.Printf("%v,%v\n", i.area(), i)
	}
	sort.Slice(squares, func(i, j int) bool { return squares[i].area() > squares[j].area() })
	fmt.Println("-----------------------")
	for _, i := range squares {
		fmt.Printf("%v,%v\n", i.area(), i)
	}
}
