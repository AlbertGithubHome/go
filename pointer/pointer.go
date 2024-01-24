package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    // 创建结构体对象
    p := Person{Name: "John", Age: 30}

    // 通过对象访问字段
    fmt.Println(p.Name)  // 输出: John
    fmt.Println(p.Age)   // 输出: 30


    // 创建结构体指针
    q := &Person{Name: "Penny", Age: 24}

    // 通过指针访问字段
    fmt.Println(q.Name)  // 输出: Penny
    fmt.Println(q.Age)   // 输出: 24
}

/*
在 Go 语言中，你可以使用点 `.` 运算符来访问结构体的字段，而不论是通过对象还是指针来引用结构体。

在上述示例中，无论是通过对象 `p` 还是通过指针 `1`，都可以使用点运算符 `.` 来访问结构体的字段。
Go 编译器会自动处理对象和指针之间的转换，使得代码更加简洁和灵活。
在使用结构体时，可以选择使用对象或指针，取决于需求和代码设计的考虑。
*/
