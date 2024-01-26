package main

import "fmt"

const (
	kConnectionStatus_None = iota
	kConnectionStatus_Connected
	kConnectionStatus_Disconnected
)

func main() {
	fmt.Println("kConnectionStatus_None:", kConnectionStatus_None)
	fmt.Println("kConnectionStatus_Connected:", kConnectionStatus_Connected)
	fmt.Println("kConnectionStatus_Disconnected:", kConnectionStatus_Disconnected)
}

/*
$ go run .
kConnectionStatus_None: 0
kConnectionStatus_Connected: 1
kConnectionStatus_Disconnected: 2
*/

/*
这段代码是Go语言中的常量定义，用于表示连接状态。它使用了iota（递增）来自动生成常量的值。具体含义如下：

1. `kConnectionStatus_None`：表示当前没有连接状态。
2. `kConnectionStatus_Connected`：表示已建立连接。
3. `kConnectionStatus_Disconnected`：表示已断开连接。

在Go语言中，`iota`是一个特殊的常量，它在每个常量组中自动递增。在这个例子中，`iota`从0开始，并且每次递增1。因此，这些常量的值将如下：

- `kConnectionStatus_None` 的值是 `0`
- `kConnectionStatus_Connected` 的值是 `1`
- `kConnectionStatus_Disconnected` 的值是 `2`

这是因为`iota`在每个新的常量声明时都会自动增加，所以每个常量都赋予了一个唯一的整数值。
*/