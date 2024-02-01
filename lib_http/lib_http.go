package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

func main() {
	// 发送 HTTP GET 请求
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应内容
	fmt.Println("Response:", resp.Status)
	fmt.Println("Body:", string(body))

	// 获取本机 IP 地址
	ip, err := getLocalIP()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Local IP:", ip)
}

// 获取本机 IP 地址
func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}


/*
$ go run .
Response: 200 OK
Body: {"ip":"1.202.147.133"}
Local IP: 10.10.49.173

在 Go 中，可以使用标准库中的 `net/http` 包来发送 HTTP 请求，
并使用 `net` 包中的函数获取本机 IP 地址。以下是一个简单的示例：

这个示例发送一个 HTTP GET 请求到 [ipify](https://www.ipify.org/) 提供的服务，
该服务返回访问者的 IP 地址信息。同时，使用 `net` 包中的 `Dial` 函数连接 Google 的 DNS 服务器，
并通过 `LocalAddr` 获取本机的 IP 地址。请注意，这只是获取本机 IP 的一种方式，
实际应用中可能需要根据实际情况选择不同的方法。
*/