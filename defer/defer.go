package main

import (
	"fmt"
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("Processing [%s] took [%s]\n", name, elapsed)
}

func main() {
	defer func(t time.Time) {
		timeTrack(t, "main")
	}(time.Now())

	time.Sleep(2 * time.Second)

	fmt.Println("sleep complete.")
}

/*
$ go run .
sleep complete.
2024/01/31 16:23:47 Processing [main] took [2.001054522s]

这段代码使用了 Go 语言中的 `defer` 语句，它的主要作用是在函数执行完成之后执行一个函数，
通常用于处理资源释放、日志记录等延迟操作。在这个具体的例子中，`defer` 用于记录函数执行的耗时。

让我们逐步解释这段代码：

1. `time.Now()`: 调用 `time.Now()` 函数获取当前的时间戳，表示函数开始执行的时间。

2. `(time.Time) { ... }(time.Now())`: 这是一个立即执行的匿名函数（自执行函数），
它接受一个 `time.Time` 类型的参数 `t`，并将 `time.Now()` 的返回值作为参数传递给这个匿名函数。
这样，我们在函数体内就可以使用 `t` 来表示函数开始执行的时间。

3. `defer`: 使用 `defer` 关键字，将匿名函数延迟执行，确保在包含 `defer` 的函数执行完毕后执行。

4. 在匿名函数体内，计算函数执行的耗时，即从开始执行到当前时刻的时间间隔，
然后通过 `log.Printf` 记录耗时信息。

总体来说，这段代码的目的是在函数执行结束时，通过 `defer` 语句自动记录函数的执行时间。
这对于性能分析、调优和日志记录等方面是很有用的，因为它避免了在每个函数调用处都手动添加记录耗时的代码。
*/