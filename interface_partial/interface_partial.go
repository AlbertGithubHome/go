package main

import "fmt"

// 定义一个接口
type Shape interface {
    Area() float64
    Perimeter() float64
}

// 定义一个矩形类型
type Rectangle struct {
    Width  float64
    Height float64
}

// 矩形类型实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    // 创建一个矩形实例
    myRectangle := Rectangle{Width: 5, Height: 3}
    myRectangle.Area();

    // Go 语言中的接口是一种契约，如果一个类型声称实现了某个接口，那么它必须实现接口中定义的所有方法。接口中的所有方法都需要在实现类型中有对应的方法
    // 尝试将矩形实例传递给接口参数（由于没有实现所有接口函数，会报编译错误Rectangle does not implement Shape (missing Perimeter method)）
    // printShapeInfo(myRectangle)
}


// 接受 Shape 接口参数的函数
func printShapeInfo(s Shape) {
    // 该函数只使用了 Shape 接口的 Area 方法
    area := s.Area()
    fmt.Println("Area:", area)

    // 无法调用 Perimeter 方法，因为矩形类型未实现该方法
    // perimeter := s.Perimeter()  // 编译错误
}
