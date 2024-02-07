package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func startDaemon() {
	goDaemon := flag.Bool("daemon", false, "run app as a daemon with -daemon=true.")
	goNohup := flag.Bool("nohup", false, "run app by nohup mode with -nohup=true.")

	flag.Parse()

	fmt.Println(os.Args)

	if *goDaemon {
		var newArgs []string

		for i := 1; i < len(os.Args); i++ {
			if !strings.Contains(os.Args[i], "-daemon") && !strings.Contains(os.Args[i], "-nohup") {
				newArgs = append(newArgs, os.Args[i])
			}
		}

		newArgs = append(newArgs, "-nohup")
		cmd := exec.Command(os.Args[0], newArgs...)

		if err := cmd.Start(); err != nil {
			fmt.Printf("start %s failed, error: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		os.Exit(0)
	} else if *goNohup {
		signal.Ignore(syscall.SIGHUP)
	}
}

func main() {
	startDaemon()
	startServer()
}

func startServer() {
	fmt.Println("Start server")
	defer func() {
		fmt.Println("Stop server")
	}()

	// wait
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

MAINLOOP:
	for {
		select {
		case <-sc:
			{
				fmt.Println("Bye~")
				break MAINLOOP
			}
		}
	}
}

/*
$ go run . -nohup
[/tmp/go-build320676713/b001/exe/daemon -nohup]
Start server
^CBye~
Stop server


$ go run . -daemon
[/tmp/go-build121262925/b001/exe/daemon -daemon]
/tmp/go-build121262925/b001/exe/daemon [PID] 443865 running...

$ kill -9 44865



startDaemon函数用于检查命令行参数，如果命令行参数包含 -daemon，则启动一个新的进程作为守护进程。
如果命令行参数包含 -nohup，则忽略 SIGHUP 信号。

1. 声明了两个命令行参数变量 `goDaemon` 和 `goNohup`，它们都是布尔型指针。这些参数允许用户在命令行中指定程序以守护进程模式运行或使用 `nohup` 模式。

2. 使用 `flag.Parse()` 解析命令行参数。这一步会解析命令行中的参数，并将其赋值给对应的变量。

3. 如果命令行参数中包含 `-daemon` 标志，则表示用户希望将程序以守护进程模式运行。在这种情况下，函数会创建一个新的进程并以守护进程模式启动它。

   - 首先，生成一个新的命令行参数列表 `newArgs`，其中移除了 `-daemon`， 添加 `-nohup` 参数。
   - 接着，创建一个新的 `exec.Cmd` 对象 `cmd`，并指定要执行的命令和参数。
   - 调用 `cmd.Start()` 方法启动新的进程。如果启动失败，则输出错误信息并退出程序。
   - 如果启动成功，则输出新进程的 PID，并退出当前进程。

4. 如果命令行参数中包含 `-nohup` 标志，则表示用户希望程序以 `nohup` 模式运行。在这种情况下，函数会忽略 `SIGHUP` 信号。

5. 如果没有 `-daemon` 和 `-nohup` 标志，则函数不执行任何操作，继续执行后续的程序逻辑。


`SIGHUP` 是 Unix-like 操作系统中的一种信号，通常由终端管理器（例如终端窗口关闭）发送给运行在前台的进程。它通常在以下情况下触发：

1. 终端窗口关闭：当用户关闭终端窗口时，会向所有在终端中运行的进程发送 `SIGHUP` 信号。

2. 登出或注销：在用户注销或登出系统时，会触发 `SIGHUP` 信号。

3. 程序异常：某些情况下，系统可能会因为程序异常而发送 `SIGHUP` 信号。

4. 手动发送：可以使用 `kill` 命令手动发送 `SIGHUP` 信号给指定的进程。
*/