package main

import (
	"fmt"
	"net"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// 发送消息给服务器
	message := "Hello, server!"
	conn.Write([]byte(message))

	// 接收服务器的回应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	response := string(buffer[:n])
	fmt.Printf("Server response: %s\n", response)
}

/*
$ go run .
Server response: Server received: Hello, server!

- 使用 `net.Dial` 函数建立 TCP 连接到服务器的地址 `"127.0.0.1:8080"`。如果连接失败，输出错误信息并终止程序。
- 使用 `defer` 关键字确保在函数结束时关闭连接。
- 创建一个消息字符串，并使用 `conn.Write` 将消息发送到服务器。
- 创建一个缓冲区，用于存储从服务器接收的数据。
- 使用 `conn.Read` 从服务器接收数据，并将数据存储到缓冲区中。如果发生错误，输出错误信息并终止程序。
- 将从服务器接收到的字节数据转换为字符串，并输出到控制台。

这个简单的 TCP 客户端连接到服务器，发送一条消息，然后接收服务器的回应并打印出来。
客户端与服务器代码一起展示了一个简单的双向通信的例子。实际应用中，你可能需要更复杂的协议和错误处理来处理各种情况。
*/