package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"unicode/utf8"
)

func codeinput() {
	hexString := "\xe5\xbe\x88\xe6\xa3\x92哈哈"
	fmt.Println(reflect.TypeOf(hexString), hexString, len(hexString))

	// 将16进制字符串转换为Unicode码点
	codePoints := []rune(hexString)
	fmt.Println(reflect.TypeOf(codePoints), codePoints, len(codePoints))

	// 将Unicode码点转换为字符串
	str := string(codePoints)
	fmt.Println(reflect.TypeOf(str), str, len(str), utf8.RuneCountInString(str))

	// 遍历码点并转换为对应的字符
	for _, codePoint := range codePoints {
		char := rune(codePoint)
		fmt.Printf("%c", char)
	}

	fmt.Print("\n")
}

func fileinput() {
	// 读取文件的全部内容
	content, err := os.ReadFile("hexfile.dat")
	if err != nil {
		log.Fatal(err)
	}

	// 使用strconv.Unquote函数去除转义字符
	unquotedText, err := strconv.Unquote(`"` + string(content) + `"`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(unquotedText)
}

func sumilateinput() {
	// 假设16进制文本为\x68\x65\x6c\x6c\x6f
	hexText := "\\xe5\\xbe\\x88\\xe6\\xa3\\x92呵呵"
	fmt.Println(reflect.TypeOf(hexText), len(hexText), hexText)

	// 使用strconv.Unquote函数去除转义字符
	unquotedText, err := strconv.Unquote(`"` + string(hexText) + `"`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(reflect.TypeOf(unquotedText), len(unquotedText), unquotedText)
}

func main() {
	codeinput()
	fmt.Println("=======================")
	fileinput()
	fmt.Println("=======================")
	sumilateinput()
}

/*
$ go run .
string 很棒哈哈 12
[]int32 [24456 26834 21704 21704] 4
string 很棒哈哈 12 4
很棒哈哈
=======================
很棒嘻嘻
=======================
string 30 \xe5\xbe\x88\xe6\xa3\x92呵呵 30
string 12 很棒呵呵


在Go中，使用len函数获取字符串的长度时，返回的是字符串的字节长度，而不是字符长度。
这意味着对于包含多字节字符的UTF-8编码字符串，每个字符可能会占用多个字节。

在Go语言中，字符串是以UTF-8编码存储的，每个字符可能占用1到4个字节。
因此，使用len()函数可以获取字符串的字节长度，而不是字符数量。

如果您需要获取字符串中Unicode字符的个数（即字符长度），可以使用utf8.RuneCountInString函数。
这个函数会返回字符串中Unicode字符的数量，而不是字节长度。


将 `\xe5\xbe\x88\xe6\xa3\x92嘻嘻` 内容写在代码中和写在文件中然后读取出来是不等价的，
写在代码中时\xe5表示一个字节，写在文件中时\xe5表示4个字节，这是不等价的。

将 \xe5\xbe\x88\xe6\xa3\x92嘻嘻 内容写在文件中然后读取出来是等价的内容是 `\\xe5\\xbe\\x88\\xe6\\xa3\\x92嘻嘻`
变成了30个字节，这时就要用的考虑转义字符的方式了，记住这种套路，在python中也是如此
unquotedText = text.encode('latin1').decode('unicode-escape')

`\\xe5` 4个字节考虑转义之后变成了 `\xe5` 也就变成了16进制的1个字节了
*/
