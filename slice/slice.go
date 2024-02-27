package main

import "fmt"

func main() {
	// 创建一个切片
	numbers := []int{1, 2, 3, 4, 5}

	// 输出切片的长度和容量
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// 对切片进行切割
	sliced := numbers[1:3]

	// 输出切片的内容
	fmt.Println("Sliced:", sliced)
}


/*
$ go run .
Length: 5, Capacity: 5
Sliced: [2 3]

- 在 Go 中，切片是对数组的一个连续片段的引用。
- 与数组相比，切片具有动态大小，且长度可以在运行时修改。
- 在示例中，我们首先创建了一个整数切片 `numbers`，其中包含了一些整数。
- 我们使用 `len` 函数获取切片的长度，并使用 `cap` 函数获取切片的容量。
- 在这里，切片的长度等于切片中元素的数量，而切片的容量是从切片的第一个元素到底层数组的最后一个元素的数量。
- 我们使用 `numbers[1:3]` 对 `numbers` 切片进行切割，产生一个新的切片 `sliced`。
- 这个切片包含了 `numbers` 切片中索引为 1 到 2 的元素，但不包含索引为 3 的元素。
- 最后，我们输出了切片 `sliced` 的内容。
*/