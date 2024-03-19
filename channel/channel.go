package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个无缓冲通道
	unbufferedChannel := make(chan int)

	// 创建一个有缓冲通道，容量为2
	bufferedChannel := make(chan string, 2)

	// 启动一个 goroutine 向无缓冲通道发送消息
	go func() {
		unbufferedChannel <- 1
		fmt.Println("Sent message to unbuffered channel")
	}()

	// 启动一个 goroutine 向有缓冲通道发送消息
	go func() {
		bufferedChannel <- "first"
		fmt.Println("Sent first message to buffered channel")
		bufferedChannel <- "second"
		fmt.Println("Sent second message to buffered channel")
	}()

	// 在主函数中读取无缓冲通道的消息，会阻塞直到有消息可以接收
	message1 := <-unbufferedChannel
	fmt.Println("Received message from unbuffered channel:", message1)

	// 在主函数中读取有缓冲通道的消息，不会阻塞
	message2 := <-bufferedChannel
	fmt.Println("Received first message from buffered channel:", message2)

	// 等待一段时间，再读取第二条消息
	time.Sleep(1 * time.Second)
	message3 := <-bufferedChannel
	fmt.Println("Received second message from buffered channel:", message3)
}

/*
$ go run .
Sent first message to buffered channel
Sent second message to buffered channel
Sent message to unbuffered channel
Received message from unbuffered channel: 1
Received first message from buffered channel: first
Received second message from buffered channel: second

通道（channel）是用来传递数据的一个数据结构，
可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
操作符 <- 用于指定通道的方向，发送或接收，如果未指定方向，则为双向通道。

声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：
ch := make(chan int)
默认情况下，通道是不带缓冲区的，发送端发送数据，同时必须有接收端相应的接收数据。

通道也可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小
ch := make(chan int, 100)

带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，
就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。

不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，
否则缓冲区一满，数据发送端就无法再发送数据了。

如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。
如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内，
如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。
接收方在有值可以接收之前会一直阻塞。
*/