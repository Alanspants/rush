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