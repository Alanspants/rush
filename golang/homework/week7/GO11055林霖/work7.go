package main

import (
	"fmt"
	"sort"
)

// 图形
type graph struct {
	area     float32
	typeName string
	graphCharacter
}

// 图形不同特征
type graphCharacter struct {
	//三角形
	// 高
	high float32
	// 底边
	bottonEdges float32
	// 长方形
	// 长度
	length float32
	// 宽度
	width float32
	//圆形
	// 半径
	radius float32
}

func NewGraph(typeName string, parms ...float32) *graph {

	gr := &graph{typeName: typeName}
	switch typeName {
	case "三角形":
		if len(parms) < 2 {
			panic("请传入三角形底边和高")
		}
		gr.high = parms[0]
		gr.bottonEdges = parms[1]
		gr.area = gr.high * gr.bottonEdges * 0.5
	case "圆形":
		if len(parms) < 1 {
			panic("请传入圆形的半径")
		}
		gr.radius = parms[0]
		gr.area = 3.14 * gr.radius * gr.radius
	case "长方形":
		if len(parms) < 2 {
			panic("请传长方形的长和宽")
		}
		gr.length = parms[0]
		gr.width = parms[1]
		gr.area = gr.length * gr.width
	default:
		panic("typeName为三角形、圆形、长方形之一")
	}

	return gr
}

func (g *graph) String() string {
	return fmt.Sprintf("%s的面积为%.2f", g.typeName, g.area)
}

type graphs []graph

func (a graphs) Len() int {
	return len(a)
}
func (a graphs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a graphs) Less(i, j int) bool { return a[i].area > a[j].area }

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	gs := make([]graph, 0, 3)

	triangle := NewGraph("三角形", 4.0, 5.0)
	round := NewGraph("圆形", 4.0)
	rectangle := NewGraph("长方形", 3.0, 4.0)
	gs = append(gs, *triangle)
	gs = append(gs, *round)
	gs = append(gs, *rectangle)

	fmt.Println(triangle)
	fmt.Println(round)
	fmt.Println(rectangle)

	// sort.Slice(gs, func(i, j int) bool { return gs[i].area > gs[j].area })
	sort.Sort(graphs(gs))
	fmt.Println("图形按面积倒序排列：")
	for _, v := range gs {
		fmt.Println(v.typeName)
	}
}
