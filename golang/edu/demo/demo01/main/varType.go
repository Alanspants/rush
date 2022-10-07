package main

import "fmt"

func main() {
	// 无符号整形，默认值都是0
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64
	fmt.Printf("u8: %d, u16: %d, u32: %d, u64: %d\n", u8, u16, u32, u64) // 默认值都为0
	u8 = 255
	u16 = 65535
	u32 = 4294967295
	u64 = 18446744073709551615
	fmt.Printf("u8: %d, u16: %d, u32: %d, u64: %d\n", u8, u16, u32, u64)

	// 整型
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	fmt.Printf("i8: %d, i16: %d, i32: %d, i64: %d\n", i8, i16, i32, i64) // 默认值都为0
	i8 = 127
	i16 = 32767
	i32 = 2147483647
	i64 = 9223372036854775807
	fmt.Printf("i8: %d, i16: %d, i32: %d, i64: %d\n", i8, i16, i32, i64)

	// int 型，取值范围32位系统为 int32，64位系统为 int64，取值相同但为不同类型
	var i int
	//i = i32 // 报错，编译不通过，类型不同
	//i = i64 // 报错，编译不通过，类型不同
	i = -9223372036854775808
	fmt.Println("i: ", i)

	// 浮点型，f32精度6位小数，f64位精度15位小数
	var f32 float32
	var f64 float64
	fmt.Printf("f32: %f, f64: %f\n", f32, f64) // 默认值都为 0.000000
	f32 = 1.12345678
	f64 = 1.12345678901234567
	fmt.Printf("f32: %v, f64: %v\n", f32, f64) // 末位四舍五入，输出：f32: 1.1234568, f64: 1.1234567890123457

	// 复数型
	var c64 complex64
	var c128 complex128
	fmt.Printf("c64: %v, c128: %v\n", c64, c128) // 实数、虚数的默认值都为0
	c64 = 1.12345678 + 1.12345678i
	c128 = 2.1234567890123456 + 2.1234567890123456i
	fmt.Printf("c64: %v, c128: %v\n", c64, c128) // 输出：c64: (1.1234568+1.1234568i), c128: (2.1234567890123457+2.1234567890123457i)

	// 字符型
	var b byte                                       // uint8 别名
	var r1, r2 rune                                  // uint16 别名
	fmt.Printf("b: %v, r1: %v, r2: %v\n", b, r1, r2) // 默认值为0
	b = 'a'
	r1 = 'b'
	r2 = '字'
	fmt.Printf("b: %v, r1: %v, r2: %v\n", b, r1, r2) // 输出：b: 97(ASCII表示的数), r1: 98(utf-8表示的数), r2: 23383 (utf-8表示的数)

	b = u8
	r1 = i32
	fmt.Printf("b: %v, r1: %v\n", b, r1) // 输出：b: 255, r1: 2147483647

	// 指针地址
	var p uintptr
	fmt.Printf("p: %v\n", p) // 默认值为0
	p = 18446744073709551615 // 64位系统最大值
	//p = 18446744073709551616 // 报错：超出最大值
	fmt.Printf("p: %v\n", p)

}
