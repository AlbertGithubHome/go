package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket connection established.")

	for {
		// 读取客户端发送的消息
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// 打印收到的消息
		fmt.Printf("Server received [%v] message: %s\n", messageType, p)

		// 回复消息给客户端
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println("Error writing message:", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("WebSocket server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting WebSocket server:", err)
	}
}


/*
$ go run .
WebSocket server is running on :8080
WebSocket connection established.
Server received [1] message: Hello, websocket!

这是一个简单的WebSocket服务器端代码，使用`github.com/gorilla/websocket`包实现。以下是对代码主要部分的解释：

- 导入了`fmt`、`github.com/gorilla/websocket`和`net/http`包。
- 定义`github.com/gorilla/websocket`包是一个用于处理WebSocket的常用库。

- `upgrader`是一个`websocket.Upgrader`对象，用于将普通HTTP连接升级为WebSocket连接。
- `CheckOrigin`函数用于检查连接的来源，这里设置为始终返回`true`，表示允许任何来源的连接。

- `handleWebSocket`函数用于处理WebSocket连接。
- 通过`upgrader.Upgrade`将HTTP连接升级为WebSocket连接，如果升级过程中发生错误，打印错误并返回。

- 在一个无限循环中，使用`conn.ReadMessage()`读取客户端发送的消息，然后打印接收到的消息。
- 接着，使用`conn.WriteMessage()`将相同的消息回复给客户端。这是一个简单的回显服务器，可以根据实际需求处理更复杂的业务逻辑。

- 在`main`函数中，使用`http.HandleFunc`设置`/ws`路径的处理函数为`handleWebSocket`。
- 然后，通过`http.ListenAndServe`启动WebSocket服务器，监听在`localhost:8080`上。如果启动过程中发生错误，打印错误信息。

这就是一个简单的WebSocket服务器端代码，可以通过连接到`ws://localhost:8080/ws`测试它。
*/