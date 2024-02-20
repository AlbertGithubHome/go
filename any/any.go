package main

import (
	"fmt"
)

func printAnyValue(value any) {
	fmt.Println("Value:", value)
}

func main() {
	// 使用空接口接收不同类型的值
	var a any
	a = 5
	printAnyValue(a) // 打印整数类型的值

	a = "Hello"
	printAnyValue(a) // 打印字符串类型的值

	a = true
	printAnyValue(a) // 打印布尔类型的值
}

/*
$ go run .
Value: 5
Value: Hello
Value: true


any 类型等价于 `interface{}`，是1.18版本加入的 定义如下
// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}
*/
