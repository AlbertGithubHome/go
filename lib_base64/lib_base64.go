package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 要编码的原始字符串
	originalString := "Hello, 世界!"

	// 编码字符串为Base64
	encodedString := base64.StdEncoding.EncodeToString([]byte(originalString))
	fmt.Println("Encoded:", encodedString)

	// 解码Base64字符串
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	decodedString := string(decodedBytes)
	fmt.Println("Decoded:", decodedString)
}

/*
$ go run .
Encoded: SGVsbG8sIOS4lueVjCE=
Decoded: Hello, 世界!

这个程序首先定义了一个原始字符串 `originalString`，然后使用 `base64.StdEncoding.EncodeToString()` 函数
将该字符串编码为Base64格式，并将结果打印出来。接着，使用 `base64.StdEncoding.DecodeString()` 函数
对该Base64字符串进行解码，并将解码后的字节序列转换回字符串，并将结果打印出来。

请注意，如果你在实际的应用程序中使用Base64编码和解码，
要确保适当地处理可能出现的错误，尤其是在解码时可能出现的错误。
*/