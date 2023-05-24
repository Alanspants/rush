package main

import "fmt"

//9*9乘法表
/*
func main() {
	for i := 1; i < 10; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d*%d=%d\t", k, i, i*k)
		}
		fmt.Print("\n")
	}
}
*/

//100以内非0随机数
/*
func main() {
	var a, t int
	t = 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rr := rand.New(r)
	for i := 0; i < 20; {
		n := rr.Intn(100)
		//出个出现0，跳出本次循环，不做后续累加次数
		if n == 0 {
			continue
		}
		fmt.Println(n)
		i++
		if n&1 == 1 {
			a += n
		} else {
			t *= n
		}
	}
	if t == 1 {
		t = 0
	}
	fmt.Printf("单数累计和:%d,偶数乘积:%d", a, t)
}
*/

//100以内的斐波那契数列
func main() {
	//1、1、2、3、5、8、13、21、34
	a := 1
	b := 1
	fmt.Println(1)
	fmt.Println(1)
	for c := 1; ; {
		c = a + b
		a = c
		b = a - b
		if c >= 100 {
			break
		}
		fmt.Println(c)
	}
}
