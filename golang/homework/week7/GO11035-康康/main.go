package main

import (
	"fmt"
	"math"
	"sort"
)


//求面积的接口
type aer interface {
	fa() float64
}
//长方形结构
type C struct {
	w float64
	h float64
}
//长方形结构体
func(c *C)fa() float64 {
	return  c.w * c.h
}
//长方形构造函数
func NewC(w ,h float64) *C {
	return  &C{w,h}
}

//圆形结构体
type Y struct {
	r float64
}
//求圆形面积的公式方法
func  (y *Y) fa() float64  {
return  y.r * math.Pi
}
//y圆形的构造函数
func NewY(i float64) *Y{
	return  &Y{i}
}

//三角形结构体S=底长×高÷2
type S struct {
	b float64 //底长
	h float64 //高
}
//三角形的求面积方法 S=底长×高÷2
func (s *S) fa() float64  {
     return s.b * s.h /2
}
//三角形构造函数
func  NewS(b,h float64) *S {
	return &S{
		b: b,
		h: h,
	}
}



//实现sort 中的interface接口 实现三个方法 修改Less方法 选择升序或者降序
type faSliceType []float64
func (x faSliceType) Len() int { return len(x) }
func (x faSliceType) Less(i, j int) bool { return x[i] > x[j]  }
func (x faSliceType) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }



func main(){
	var faSlice = make([]float64,0,3)
	//圆形求面积
	y := NewY(132)
	var a aer = y
	fmt.Println(a.fa())
	faSlice=append(faSlice,a.fa())
	//长方形求面积
	c := NewC(10,29)
	a  = c
	fmt.Println(a.fa())
	faSlice=append(faSlice,a.fa())
	//三角形求面积
	s := NewS(13,44)
	a = s
	fmt.Println(a.fa())
	faSlice=append(faSlice,a.fa())
    fmt.Println(faSlice)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	//降序
	fmt.Println("降序～～～～～～实现接口，实现三个方法")
	sort.Sort(faSliceType(faSlice))
	fmt.Println(faSlice)

	//升序
	fmt.Println("升序～～～～～～使用SliceStable 使用高阶函数")
	sort.SliceStable(faSlice, func(i, j int) bool {
		return  faSlice[i] < faSlice[j]
	})
	fmt.Println(faSlice)

}
