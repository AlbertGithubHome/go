package main

import "fmt"

func main() {

	// 使用单引号表示一个 rune 类型值
	var r rune = '你'

	// 打印这个 rune 类型值
	fmt.Println(r)
    fmt.Printf("%d, %c, %U\n", r, r, r)

	// 将 rune 类型转换为字符串类型
	s := string(r)
	fmt.Println(s)

    fmt.Println("==============================")

	// 定义一个包含 Unicode 字符的字符串
	str := "Hello, 世界"

	// 将字符串转换为 rune 切片
	runes := []rune(str)

	// 遍历 rune 切片并打印每个字符的 Unicode 值和字符本身
	fmt.Println("Unicode code points and characters:")
	for i, r := range runes {
		fmt.Printf("%d, Unicode: %U, Character: %c\n", i, r, r)
	}
}

/*
$ go run .
20320
20320, 你, U+4F60
你
==============================
Unicode code points and characters:
0, Unicode: U+0048, Character: H
1, Unicode: U+0065, Character: e
2, Unicode: U+006C, Character: l
3, Unicode: U+006C, Character: l
4, Unicode: U+006F, Character: o
5, Unicode: U+002C, Character: ,
6, Unicode: U+0020, Character:
7, Unicode: U+4E16, Character: 世
8, Unicode: U+754C, Character: 界

这段代码演示了在 Go 中如何使用 `rune` 类型处理 Unicode 字符。

1. 首先，我们定义了一个 `rune` 类型的变量 `r`，并初始化为一个 Unicode 字符 `'你'`。
2. 然后，我们打印这个 `rune` 类型的值，以及它的 Unicode 编码、字符本身。
3. 接着，我们将 `rune` 类型转换为字符串类型，并打印出来。
4. 最后，我们定义了一个包含 Unicode 字符的字符串 `str`，将其转换为 `rune` 切片 `runes`，
  并使用 `range` 循环遍历该切片，打印每个字符的索引、Unicode 值和字符本身。

整体来说，这段代码演示了如何在 Go 中使用 `rune` 类型来处理 Unicode 字符，包括初始化、转换和遍历等操作。


字符是 UTF-8 编码的 Unicode 字符，Unicode 为每一个字符而非字形定义唯一的码值（即一个整数），

例如 字符a 在 unicode 字符表是第 97 个字符，所以其对应的数值就是 97，

也就是说对于Go语言处理字符时，97 和 a 都是指的是字符a，而 Go 语言将使用数值指代字符时，将这样的数值称呼为 rune 类型。

rune类型是 Unicode 字符类型，和 int32 类型等价，通常用于表示一个 Unicode 码点。

rune 和 int32 可以互换使用。 一个Unicode代码点通常由”U+”和一个以十六进制表示法表示的整数表示，例如英文字母’A’的Unicode代码点为”U+0041”。

在 Go 中，rune 类型是用于表示 Unicode 码点的类型。Unicode 是一种标准，用于为世界上各种语言和符号分配唯一的数字编码，以便它们可以在计算机中存储和处理。

在 Go 中，rune 类型实际上是一个 int32 类型的别名。它可以用于表示任何 Unicode 码点，并提供了一些有用的函数和方法，用于处理字符串和 Unicode 编码。


需要注意的是，在处理字符串和 Unicode 编码时，我们需要特别注意编码和解码的方式，以避免出现错误或者不正确的结果。
同时，我们也要了解各种字符集和编码标准之间的差异，以便在处理多语言和多文化环境下的应用程序时保持最佳实践。

此外rune类型的值需要由单引号”‘“包裹，不过我们还可以用另外几种方式表示:

1. 直接使用单引号表示一个 Unicode 字符，例如：`var r rune = '你'`。
2. 使用十六进制表示 Unicode 码点，例如：`var r rune = '\u4F60'`。
3. 使用十进制表示 Unicode 码点，例如：`var r rune = 20320`。
4. 使用转义字符表示 Unicode 字符，例如：`var r rune = '\n'` 表示换行符。

*/