package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建两个通道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个 goroutine，分别向两个通道发送消息
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from channel 2"
	}()

MAINLOOP:
	for {
		// 使用 select 语句等待多个通信操作
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout: No communication within 3 seconds")
			break MAINLOOP
		}
	}

}

/*
$ go run .
Message from channel 2
Message from channel 1
Timeout: No communication within 3 seconds


在 Go 语言中，select 语句用于处理一组通信操作，可以与 case 语句搭配使用。
select 语句使得程序可以等待多个通信操作中的任意一个完成，并执行相应的代码块。

在上述例子中，select 语句用于等待多个通信操作，包括从 ch1 接收消息、从 ch2 接收消息以及一个定时器（time.After）的超时操作。
select 会等待其中一个 case 操作完成，然后执行相应的代码块。如果所有 case 都未完成，可以使用 default 语句执行默认的操作。

这样的 select 语句非常适用于多个通道之间的非阻塞并发操作，使得程序能够响应多个通信操作的完成情况。
*/