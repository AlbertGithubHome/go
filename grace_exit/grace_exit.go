package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
这段代码使用了 Go 中的信号（Signal）和通道（Channel）机制，用于监听操作系统发出的信号（如 SIGINT、SIGTERM、SIGQUIT）并在接收到信号时执行相应的操作。

1. `sc := make(chan os.Signal, 1)`：创建一个带有缓冲区大小为 1 的信号通道 `sc`，用于接收操作系统发送的信号。

2. `signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)`：使用 `signal.Notify` 函数注册 `sc` 通道，以便接收指定的信号（SIGINT、SIGTERM、SIGQUIT）。

3. `for range sc`：使用 `for range` 语法监听信号通道，当通道接收到信号时，执行循环体内的操作。

4. `fmt.Println("Bye~")`：当收到信号时，打印 "Bye~" 消息。

5. `break MAINLOOP`：跳出 `MAINLOOP` 标签所标记的循环，即终止循环。

整体来说，这段代码的作用是等待并监听指定的信号（SIGINT、SIGTERM、SIGQUIT），一旦接收到信号，就打印 "Bye~" 消息并跳出循环，达到程序优雅退出的效果。
*/

func main() {
	fmt.Println("Hello,World!")

	// wait
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

MAINLOOP:
	for range sc {
		fmt.Println("Bye~")
		break MAINLOOP
	}
}

/*
在上述的简单例子中，使用 `MAINLOOP` 标签并不是必须的，因为只有一个循环。标签主要在多层嵌套循环时发挥作用，用于指定 `break` 或 `continue` 跳出或继续的是哪个循环。

在单层循环的情况下，去掉标签并直接使用 `break` 是更常见和简洁的写法。然而，在多层循环中，标签可以帮助准确定位到底是哪个循环被中断或继续，增强了代码的可读性。

例如：

```go
outerLoop:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Printf("%d-%d ", i, j)
        if i == 1 && j == 1 {
            break outerLoop // 中断outerLoop标签所标记的循环
        }
    }
}
```

在这个例子中，`break outerLoop` 将中断 `outerLoop` 标签所标记的外层循环，而不是内层循环。

总体来说，使用标签的好处主要在于提高代码的清晰度和可读性，特别是在存在多层循环的情况下。在单层循环的情况下，去掉标签是更简洁的写法。



`syscall.SIGINT`、`syscall.SIGTERM` 和 `syscall.SIGQUIT` 是在 Unix-like 操作系统上定义的信号，它们用于与正在运行的程序进行通信，例如请求程序终止等。

这些信号可以通过终端键盘输入或其他方式触发。以下是它们的含义和触发方式：

1. **SIGINT（Interrupt）：**
   - **含义：** 通常由用户在终端上按下 Ctrl+C 键触发。用于中断当前进程的运行。
   - **触发方式：** Ctrl+C 键盘组合。

2. **SIGTERM（Termination）：**
   - **含义：** 表示终止信号，通常用于请求程序正常终止。程序收到 SIGTERM 时，可以进行清理操作后退出。
   - **触发方式：** 通过 `kill` 命令或操作系统的其他工具发送。

3. **SIGQUIT（Quit）：**
   - **含义：** 类似于 SIGTERM，但同时生成核心转储文件，用于调试程序。通常用于请求程序终止，并生成调试信息。
   - **触发方式：** 通常由用户在终端上按下 Ctrl+\ 键触发。

这些信号提供了一种与正在运行的程序进行交互的机制。上述代码中，通过使用 `signal.Notify` 注册了对这些信号的监听，
当程序接收到这些信号时，会执行相应的操作。例如，当收到 `syscall.SIGINT` 信号时，程序会打印 "Bye~" 消息并退出。
*/
