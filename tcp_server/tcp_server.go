package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 处理连接，这里简单回显收到的消息
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			// 判断是否为 EOF 错误
			if err.Error() == "EOF" {
				fmt.Println("Client closed the connection.")
				return
			}
			fmt.Println("Error reading:", err)
			return
		}

		message := string(buffer[:n])
		fmt.Printf("Received message: %s\n", message)

		// 回显消息给客户端
		conn.Write([]byte("Server received: " + message))
	}
}

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	// 接受连接并处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

/*
$ go run .
Server is listening on port 8080
Received message: Hello, server!
Client closed the connection.

handleConnection 函数接受一个 net.Conn 类型的参数，代表与客户端的连接。
使用 defer 关键字确保在函数结束时关闭连接。
使用一个无限循环读取客户端发送的消息，将消息打印到控制台，并回显相同的消息给客户端。

main 函数开始监听 TCP 端口 127.0.0.1:8080。
在无限循环中，通过 listener.Accept() 接受客户端的连接请求。一旦有连接请求，
就创建一个新的协程（goroutine）调用 handleConnection 函数来处理连接。
使用 defer 关键字确保在程序退出时关闭监听器。

这个简单的 TCP 服务器会一直监听指定端口，每当有客户端连接时，创建一个新的协程处理连接。
在连接被接受后，服务器会不断读取客户端发送的消息并进行回显。这个示例主要用于演示 TCP 服务器的基本结构，
实际中可能需要更多的错误处理和更复杂的业务逻辑。
*/
