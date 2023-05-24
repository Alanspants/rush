//1、实现对圆形、三角形、长方形求面积
//2、利用第1题，构造3个以上图形，至少圆形、三角形、矩形各有一个，对上题的图形按照面积降序排列

package main

import (
    "fmt"
    "sort"
)

type Areaer interface {
    area(float64, float64) float64
}

type Circle struct {
    Pi float64
    Radial  float64
    Area    float64
}

func (*Circle) area(pi,r float64) float64 {
    Area :=  pi*r*r
    return Area
}

type Triangle struct {
    bottom,high,Area float64
}

func (*Triangle) area(b,h float64) float64 {
    Area :=  b*h/2
    return Area
}

type Rectangle struct {
    length,wide,Area float64
}

func (*Rectangle) area(l,w float64) float64 {
    Area := l*w
    return Area
}

func garea(a Areaer,x float64,y float64) {
    a.area(x,y)
}

//在接口之上封装普通函数,接口为参数,因为这三种图形结构体都实现了接口,因此都满足该接口,调用接口参数的方法,即是调用对应结构体实例的方法,实现了多态

func main() {
    c := new(Circle)
    c.Pi = 3.14
    c.Radial = 10
    c.Area = c.area(c.Pi,c.Radial)
    fmt.Printf("圆形面积为: %v\n",c.Area)
    t :=new(Triangle)
    t.bottom = 10
    t.high = 5.5
    t.Area = t.area(t.bottom,t.high)
    fmt.Printf("三角形面积为: %v\n", t.Area)
    r :=new(Rectangle)
    r.length = 15.3
    r.wide = 9.9
    r.Area = r.area(r.length,r.wide)
    fmt.Printf("矩形面积为: %v\n", r.Area)
    
    areas := make([]float64,0)
    fmt.Printf("面积排序: %v",areas)
    areas = append(areas,c.Area)
    areas = append(areas,t.Area)
    areas = append(areas,r.Area)
    sort.Float64s(areas)
    fmt.Printf("面积排序: %v\n",areas)
    

}


