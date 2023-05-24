package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//第一题，打印九九乘法表
	ChenFaBiao()

	//第二题打印100内20个随机数。。。
	RandSum(100)

	//第三题打印100内的斐波那契数列
	FeiBoNaqie()
}

func ChenFaBiao(){
	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			if i*j < 10{
				fmt.Printf("%d*%d = %d  ",j,i,i*j)
			}else {
				fmt.Printf("%d*%d = %d ",j,i,i*j)
			}
			if i == j {
				fmt.Printf("\n")
			}
		}
	}
}

func RandSum(n int){
	ln := make([]int,0,20)
	for i:=0;i<20;i++{
		n := rand.Intn(n)
		if n == 0{
			i--
		}else {
			ln = append(ln,n)
		}
	}
	fmt.Println(ln,len(ln))
	sum,res := 0,1
	for i := 0;i < len(ln);i++{
		if i % 2 == 0 {
			sum+=ln[i]
		}else {
			res*=ln[i]
		}
	}
	fmt.Println(sum,res)
}

//第三题打印斐波那契数列1，1，2，3，5，8，13，21，34，55。。。
func FeiBoNaqie(){
	//if n==1 || n==2{
	//	return 1
	//}else {
	//	return FeiBoNaqie(n-1)+FeiBoNaqie(n-2)
	//}

	a,b := 1,1
	for {
		a,b = b,a+b
		if a < 100 {
			fmt.Println(a)
		}else {
			break
		}
	}

}
