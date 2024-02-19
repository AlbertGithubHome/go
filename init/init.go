package main

import (
	"fmt"

	"init/foo" // 导入自定义包
)

func main() {
    fmt.Println("Hello, world!")
    fmt.Println("The value of x:", foo.GetX())
}

/*
$ go run .
foo package is initialized
Hello, world!
The value of x: 18

在这个示例中，我们定义了一个自定义包 `foo`，并在其中定义了一个全局变量 `x`，
以及一个初始化函数 `init`。在 `init` 函数中，我们给 `x` 赋值为 `18`，
并打印了初始化消息。在 `main` 包中，我们导入了自定义包 `foo`，
并调用了其中的 `GetX` 函数来获取 `x` 的值。当程序运行时，
会首先执行 `foo` 包的 `init` 函数，然后再执行 `main` 函数。

`init` 函数在包被导入时自动执行，用于执行一些初始化操作，
例如初始化全局变量、注册资源等。在实际开发中，可以利用 `init` 函数
来执行一些必要的初始化工作，以确保包在被使用之前处于正确的状态。
*/