

# 基础

## 入门

### hello world

```go
package main

import "fmt"

func main() {
	fmt.Println("hello, world!")
}
```

### 编译执行

```shell
$ go build hello.go
$ ./hello
hello, world!

$ go run hello.go
hello, world!
```

区别：

1. 可以放到没有go环境下运行
2. 如果go run，需要别的机器也有go环境
3. 编译器会将程序运行以来的库文件包含在可执行文件中，所以可执行文件变大了很多

### 注意事项

1. 执行入口：main()函数
2. 严格区分大小写
3. 不需要分号
4. 定义的变量或import的pkg没有用到是不能compile pass的

## 语法

### 包

包名和导入路径的最后一个元素一致，例如，“math/rand”包中的源码均以rand语句开始

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

如果一个名字以大写字母开头，那么就是已导出的。

```go
correct：fmt.Println(math.Pi)
wrong: fmt.Println(math.pi)
```

### 函数

函数类型在变量名之后。

当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

函数可以返回任意数量的结果。

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

OUTPUT:
world hello
```

return值可以被命名。

并且可以在函数的结尾用return（不加返回值）直接返回，但是最好用于短函数中，长函数会影响可读性。

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

OUTPUT:
7 10
```

### 变量

var声明的变量，类型在末尾，同时同类型的变量也可以生成一大串。

```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

OUTPUT:
0 false false false
```

同一个类型的变量可以在一行中初始化。

不同类型的变量，可以不加类型并初始化。

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

OUTPUT:
1 2 true false no!
```

简洁赋值语句 := 可以在类型明确的地方代替 = 使用。

函数外的每个语句都必须以关键字开始（var，func等等），因此 := 结构不能再在函数外使用。

```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

OUTPUT:
1 2 3 true false no!
```

数值类型初始化为 `0`，

布尔类型初始化为 `false`，

字符串初始化为 `""`（空字符串）。

### 常量

常量可以是字符、字符串、bool、数值。

常量不能用 := 声明。

### 整数类型

int8 有符号整型：范围是 -128 ~ 127

uint8 无符号整型：范围是 0 ~ 255

byte：0 ~ 255

---

<strong style="color:red;">为什么是-128到127，而不是-127到127</strong> 

源码：符号位 + 真值绝对值，0正数，1负数

反码：正数反码不变，负数的反码 -> 符号位不变，其他依次取反

补码：正数补码不变，负数的补码 -> 反码 + 1

+127的补码是0111 1111，+128就是0000 0000，刚好循环到0，所以正值最大是127

-127的补码是1000 0001，-128就是1000 0000，所以-0=-128=1000 0000

---

如何查看变量的类型。

```go
var n1 = 100
fmt.Printf("n1's type: %T", n1)

OUTPUT: n1's type: int
```

golong程序中整型变量使用时，遵守保小不保大的原则。

### 字符类型

用byte表示，0 ~ 255的ascii/utf-8(unicode)

当直接输出byte值时，就是输出了对应字符的ascii值，如果希望输出对应字符，需要使用格式化输出。

```go
var c1 byte = 'a'
var c2 byte = 'A'

fmt.Println("c1=", c1)
fmt.Println("c2=", c2)

fmt.Println("c1=%c, c2=%c", c1, c2)

OUTPUT:
c1=97
c2=65
c1=a, c2=A
```

如果超出255，可以先用int存

```go
var c3 int = '北'
fmt.Println("c3=%c, c3 ascii=%d", c3, c3)

OUTPUT:
c3=北, c3 ascii=21271
```

utf-8其中包含了ascii码值，英文一个字节，汉字三个字节

### 布尔类型

不可以用0或非0的数值代替false和true

### 字符串类型

1. 使用utf-8
2. golang中的字符串是不可变的
3. golong中可以使用反引号（``)将字符串以原生形式输出
4. 字符串拼接

```go
var str := "hello" + "world"
str += "haha"

OUTPUT:
helloworldhaha
```

基本类型 -> string类型

```go
var num1 int = 99
str := fmt.Sprintf("%d", num1)
str := strcov.FormatInt(num1, 10)

var num2 float64 = 23.456
str := fmt.Sprintf("%f", num2)
str := strconv.FormatFloat(num2, 'f', 10, 64)

var num3 int = 4567
str := strconv.Itoa(num3)
```

string类型 -> 基本类型

```go
var str string = "true"
var b bool
b, _ = strconv.ParseBool(str)
fmt.Printf("b type: %T, b value: %v", b)

OUTPUT:
b type: bool, b value: true
```

如果一个不能被转成目标的数，会被转成0

```go
var str string = "hello"
var n3 int = 0
n3, _ = strconv.ParseInt(str, 10, 64)
fmt.Println(n3)

OUTPUT:
0
```

### 指针

&变量名 -> 查看一个变量的地址

```go
var i int = 10
fmt.Println("i address: ", &i)

OUTPUT:
0*0cc01c0
```

指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值。比如：var ptr **int = &num*

```go
var i int = 10
// ptr是一个指针变量
// ptr本身的值是&i
var ptr *int = &i

// ptr的值 -> 存的是i的地址
fmt.Printf("ptr=%v\n", ptr)
// ptr自身的地址
fmt.Printf("&ptr=%v\n", ptr)
// ptr的值（i的地址）所指向的变量值
fmt.Printf("*ptr=%v\n", ptr)
```

 ### 运算符

除法：/ ，直接去掉小数部分，保留整数部分

如果希望保留小数部分，则需要有浮点数参与计算

```go
fmt.Println(10 / 4)
var n2 float32 = 10.0 / 4

OUTPUT:
2
2.5
```

模：%

公式：a % b = a - a / b * b

### 逻辑运算符

短路运算符：&& ｜｜

如果第一个条件符合，第二个条件不会判断

### 源码，反码，补码

1. 二进制最高位的符号位：0 -> 正数， 1 -> 负数
2. 正数的源码，补码，反码都一样
3. 负数的反码 = 源码符号位不便，其他位取反
4. 负数的补码 = 负数的反码 + 1
5. 0的反码补码都是0
6. 在计算机运算的时候，都是以补码的方式来运算的

## 代码

### 闭包

```go
func AddUpper() func (int) int {
  var n int = 10
  var str = "hello"
  return func (x int) int {
    n = n + x
    return n
  }
}

func main() {
  f := AddUpper()
  fmt.Println(f(1)) //11
  fmt.Println(f(2)) //13
  fmt.Println(f(3)) //16
}
```

闭包中，函数整体相当于一个类，其中定义的变量相当于成员变量，只会在这个函数被初始化的时候执行一次，函数内自带的函数相当于成员方法，每次会在成员变量初始化的基础上反复执行。 

### defer

```go
func sum(n1 int, n2 int) int {
  defer fmt.Println("ok1 n1=", n1)// 第三个输出
  defer fmt.Println("ok1 n2=", n2)// 第二个输出
  
  res: = n1 + n2
  fmt.Println("okk3 res=", res)		// 第一个输出
}
```

当执行到defer时，暂时不会执行，会将defer后面的语句压入到独立的栈（defer栈）

当函数执行完毕后，再从defer栈，按照先入后出的顺序出栈，执行

```go
func sum(n1 int, n2 int) int {
  defer fmt.Println("ok1 n1=", n1)// 第三个输出 
  defer fmt.Println("ok1 n2=", n2)// 第二个输出
  
  n1++
  n2++
  res: = n1 + n2
  fmt.Println("okk3 res=", res)		// 第一个输出
}
```

下侧n1和n2的变化，对print没有影响，因为在++之前，相关变量就已经被压入栈了。

```go
file = openfile(filename)
defer file.close()

connect = openDatabase()
defer connect.close()
```

