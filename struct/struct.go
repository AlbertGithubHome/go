package main

import "fmt"

// 定义一个结构体类型
type Person struct {
    Name    string
    Age     int
    Address string
}

type Point struct {
    X, Y int
}

// Go语言有一个特性让我们只声明一个成员对应的数据类型，而不指名成员的名字，这类成员就叫匿名成员。
type Circle struct {
    Point  // 匿名字段，struct
    Radius int
}
type Wheel struct {
    Circle  // 匿名字段，struct
    Spokes int
}

func testEmbed() {
    w := Wheel{Circle{Point{8, 8}, 5}, 20}
    fmt.Printf("%#v\n", w)

    //在右边的注释中给出的显式形式访问这些叶子成员的语法依然有效,因此匿名成员并不是真的无法访问了。
    //其中匿名成员Circle和Point都有自己的名字——就是命名的类型名字——但是这些名字在点操作符中是可选的。

    //我们在访问子成员的时候可以忽略任何匿名成员部分。
    var w2 Wheel
    w2.Circle.Point.X = 8
    w2.Circle.Point.Y = 8
    w2.Circle.Radius = 5
    w2.Spokes = 20

    //结构体字面值并没有简短表示匿名成员的语法, 因此下面的语句都不能编译通过
    //w = Wheel{8, 8, 5, 20}   // compile error: unknown fields
    //w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields
}

func main() {
    // 创建一个结构体实例并初始化
    p1 := Person{Name: "Alice", Age: 25, Address: "110 Main St"}

    // 访问结构体字段并打印
    fmt.Println("Name:", p1.Name)
    fmt.Println("Age:", p1.Age)
    fmt.Println("Address:", p1.Address)

    // 修改结构体字段的值
    p1.Age = 30
    p1.Address = "120 Elm St"

    // 打印修改后的结构体字段值
    fmt.Println("Modified Age:", p1.Age)
    fmt.Println("Modified Address:", p1.Address)

    // 注意事项：
    // 1. 结构体是一种自定义的数据类型，可以包含多个字段。
    // 2. 使用点号（.）访问结构体字段。
    // 3. 可以通过赋值语句修改结构体字段的值。
    // 4. 结构体的字段名是大小写敏感的。

    testEmbed()
}

/*
$ go run .
Name: Alice
Age: 25
Address: 110 Main St
Modified Age: 30
Modified Address: 120 Elm St
main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}

一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身（该限制同样适应于数组）。
但是S类型的结构体可以包含*S指针类型的成员，这可以让我们创建递归的数据结构，比如链表和树结构等。

Go语言有一个特性让我们只声明一个成员对应的数据类型，而不指名成员的名字，这类成员就叫匿名成员。

当使用 %#v 格式化指示符时，它会以 Go 语法格式输出值，包括类型信息和内部结构。
对于复杂的数据类型，这种格式化方式非常有用，可以方便地查看值的详细结构
*/