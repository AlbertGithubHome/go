package main

import (
	"os"

	"github.com/withmandala/go-log"
)

func main() {
	// 创建一个新的日志记录器
	logger := log.New(os.Stdout).WithColor().WithDebug()

	// 打印一些信息
	logger.Debug("这是一条调试信息")
	logger.Info("这是一条普通信息")
	logger.Warn("这是一条警告信息")
	logger.Error("这是一条错误信息")
	logger.Fatal("这是一条致命错误信息")
}


/*
$ go run .
[DEBUG] 2024/02/08 14:45:09 main.main:lib_log.go:14 这是一条调试信息
[INFO]  2024/02/08 14:45:09 这是一条普通信息
[WARN]  2024/02/08 14:45:09 这是一条警告信息
[ERROR] 2024/02/08 14:45:09 main.main:lib_log.go:17 这是一条错误信息
[FATAL] 2024/02/08 14:48:22 main.main:lib_log.go:18 这是一条致命错误信息
exit status 1


这段代码是一个简单的Go程序，使用了 `github.com/withmandala/go-log` 包来记录不同级别的日志信息。

1. 在导入部分，我们导入了 `os` 包和 `github.com/withmandala/go-log` 包。

2. 在 `main()` 函数中：

   - 我们使用 `log.New(os.Stdout).WithColor().WithDebug()` 创建了一个新的日志记录器实例。
   - `os.Stdout` 作为参数传递给 `log.New()` 函数，这意味着日志消息将被输出到标准输出流。
   - `.WithColor()` 方法用于启用日志消息的彩色输出，`.WithDebug()` 方法用于设置日志级别为调试级别。

   - 然后，我们使用不同级别的方法打印了一些日志信息：
     - `logger.Debug()` 打印调试信息。
     - `logger.Info()` 打印普通信息。
     - `logger.Warn()` 打印警告信息。
     - `logger.Error()` 打印错误信息。
     - `logger.Fatal()` 打印致命错误信息，并终止程序的执行。

这段代码演示了如何使用 `github.com/withmandala/go-log` 包记录不同级别的日志信息，并将其输出到标准输出流。
*/