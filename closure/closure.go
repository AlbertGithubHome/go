package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos := adder()
	for i := 0; i < 5; i++ {
		fmt.Println("Positive sum:", pos(i))
	}
}

/*
$ go run .
Positive sum: 0
Positive sum: 1
Positive sum: 3
Positive sum: 6
Positive sum: 10

这个示例中，我们只调用了 adder 函数一次，并得到了一个闭包函数 pos。
然后，我们循环调用 pos 函数，传入不同的整数，观察累加的结果。

Go语言是支持闭包的，这里只是简单地讲一下在Go语言中闭包是如何实现的。

匿名函数是无需定义标示符(函数名)的函数；而闭包是指能够访问自由变量的函数。

换句话说，定义在闭包中的函数可以”记忆”它被创建时候的环境。闭包函数=匿名函数＋环境变量。

总的来说： 在 Go 中，闭包函数是一种特殊的匿名函数，可以访问其外部作用域中的变量，并在函数内部持久化存储它们。

通过使用闭包函数，我们可以轻松地实现很多高级编程技巧，例如延迟计算、惰性求值、缓存等。
需要注意的是，在使用闭包函数时，我们需要特别注意变量的生命周期和作用域，以避免出现内存泄漏和其他错误。
*/