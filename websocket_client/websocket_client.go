package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	url := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("Error connecting to WebSocket:", err)
		return
	}
	defer conn.Close()

	done := make(chan struct{})

    msg := []byte("Hello, websocket!");
	fmt.Printf("Client send message: %s\n", msg)
	conn.WriteMessage(websocket.TextMessage, msg)

	// 启动一个 goroutine 用于读取消息
	go func() {
		defer close(done)
		for {

			messageType, p, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}
			fmt.Printf("Client received message: %s\n", p)

			// 在这里可以根据接收到的消息进行处理
			// 例如，如果收到特定的消息，可以关闭连接或执行其他操作
			_ = messageType
		}
	}()

	// 主 goroutine 用于发送消息
	for {
		select {
		case <-interrupt:
			fmt.Println("Interrupt signal received, closing connection.")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("Error closing connection:", err)
			}
			return
		}
	}
}

/*
$ go run .
Client send message: Hello, websocket!
Client received message: Hello, websocket!


- 导入了`fmt`、`github.com/gorilla/websocket`包和`os`包，用于处理输入输出和信号。`os/signal`包用于捕获中断信号。

- 在主函数中，首先创建一个用于捕获中断信号的通道`interrupt`，然后使用`signal.Notify`订阅中断信号。
- 接着，定义WebSocket服务器的地址`url`，并使用`websocket.DefaultDialer.Dial`连接到服务器。

- 启动两个goroutine：
   - 一个用于读取服务器发送的消息。
   - 主goroutine用于等待中断信号。

- 在主goroutine中，通过`select`语句监听中断信号。一旦接收到中断信号，就会关闭WebSocket连接，结束程序。

这个WebSocket客户端通过在两个goroutine中分别处理消息读取和中断信号，实现了异步处理服务器消息的功能。
当接收到中断信号时，客户端会关闭连接并退出。
*/
