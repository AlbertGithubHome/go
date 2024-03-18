package main

import (
	"fmt"
	"time"
)

func recoverFunc() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from:", r)
	}
}

func worker() {
	defer recoverFunc()

	fmt.Println("Worker started...")
	time.Sleep(2 * time.Second)

	// 模拟出现 panic
	panic("Something went wrong in worker!")
}

func main() {
	fmt.Println("Starting main...")

	// 启动一个 goroutine 执行 worker 函数
	go worker()

	// 主 goroutine 继续执行其他任务
	fmt.Println("Main goroutine continues...")

	// 等待一段时间，确保 worker goroutine 有足够的时间执行
	time.Sleep(3 * time.Second)

	fmt.Println("Main goroutine exits.")
}

/*
$ go run .
Starting main...
Main goroutine continues...
Worker started...
Recovered from: Something went wrong in worker!
Main goroutine exits.

这是一个演示了在 goroutine 中触发 panic 的示例，但程序不会被关闭

在这个示例中，我们定义了一个 `worker` 函数，其中包含了一个 `panic` 语句。
在 `main` 函数中，我们通过 `go` 关键字启动了一个 goroutine 来执行 `worker` 函数，
这样 `worker` 函数的执行不会阻塞 `main` 函数的执行。

即使在 `worker` 函数中触发了 panic，由于我们在 `worker` 函数中使用了 `recover` 函数，
因此程序不会被关闭。相反，panic 会被捕获，并在 `recoverFunc` 函数中进行处理。
出现这种情况时，我们简单地打印了恢复的信息。

通常情况下，`recover` 函数应该放在 `defer` 语句中调用。
这样做的好处是，无论函数体内部发生了什么，`defer` 中的代码都会在函数退出时执行，
从而可以在出现 panic 的情况下进行恢复操作。

将 `recover` 放在 `defer` 中的原因是为了确保它能够捕获到函数体内部的 panic。
如果 `recover` 不在 `defer` 中调用，而是直接在函数体内调用，
那么在函数体内部的 panic 将无法被 `recover` 捕获，从而导致程序终止。

如果将 defer recoverFunc() 语句注释掉，将会输出下面的结果
$ go run .
Starting main...
Main goroutine continues...
Worker started...
panic: Something went wrong in worker!

goroutine 18 [running]:
main.worker()
        /home/shz/WorkSpace/gospace/gopractice/panic_recover/panic_recover.go:21 +0x6f
created by main.main
        /home/shz/WorkSpace/gospace/gopractice/panic_recover/panic_recover.go:28 +0x65
exit status 2
*/
