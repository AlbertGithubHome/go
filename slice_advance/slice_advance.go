package main

import "fmt"

func main() {
    var slice []int // 声明一个切片变量，值为 nil，长度和容量为 0

    // 添加元素前需要先初始化切片
    // slice = make([]int, 0) // 使用 make 函数初始化切片

    // 直接向切片中添加元素
    slice = append(slice, 1)
    slice = append(slice, 2)

    fmt.Println(slice, len(slice), cap(slice))

    // 声明一个整型切片
    var numbers []int

    // 循环向 numbers 切片中添加 10 个数
    for i := 0; i < 10; i++ {
        numbers = append(numbers, i)
        // 打印输出切片的长度、容量和指针变化，使用函数 len() 查看切片拥有的元素个数，使用函数 cap() 查看切片的容量情况
        fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
    }

    testCopy()
}

func testCopy() {
    src := []int{1, 2, 3, 4, 5}
    dst1 := make([]int, 3)
    dst2 := make([]int, 3)

    copy(dst1, src)
    copy(dst2, src)

    fmt.Printf("dst1: %v dst2: %v\n", dst1, dst2)
}

/*
[1 2] 2 2
len: 1  cap: 1 pointer: 0xc0000b4040
len: 2  cap: 2 pointer: 0xc0000b4050
len: 3  cap: 4 pointer: 0xc0000b6020
len: 4  cap: 4 pointer: 0xc0000b6020
len: 5  cap: 8 pointer: 0xc0000b8040
len: 6  cap: 8 pointer: 0xc0000b8040
len: 7  cap: 8 pointer: 0xc0000b8040
len: 8  cap: 8 pointer: 0xc0000b8040
len: 9  cap: 16 pointer: 0xc0000ba000
len: 10  cap: 16 pointer: 0xc0000ba000
dst1: [1 2 3] dst2: [1 2 3]

// The append built-in function appends elements to the end of a slice. If
// it has sufficient capacity, the destination is resliced to accommodate the
// new elements. If it does not, a new underlying array will be allocated.
// Append returns the updated slice. It is therefore necessary to store the
// result of append, often in the variable holding the slice itself:
//
//	slice = append(slice, elem1, elem2)
//	slice = append(slice, anotherSlice...)
//
// As a special case, it is legal to append a string to a byte slice, like this:
//
//	slice = append([]byte("hello "), "world"...)

以上是slice的append函数注释，当没有足够的空间时append会分配新的数组，
这就是为什么不使用make函数直接append元素而不会报错的原因
而 map 和 channel 只定义不使用make分配就会panic

通过测试可以发现，每次append添加元素后扩展空间的大小遵循规律1,2,4,8，每次扩展后容量翻倍
*/