package main

import "fmt"

// 定义一个接口
type Quacker interface {
    Quack()
}

// 定义一个结构体
type Duck struct {
    Name string
}

// Duck 类型实现 Quacker 接口的 Quack 方法
func (d Duck) Quack() {
    fmt.Println(d.Name, "quacks")
}

// 函数接受一个 Quacker 接口参数
func makeItQuack(q Quacker) {
    q.Quack()
}

func main() {
    // 创建一个 Duck 实例
    myDuck := Duck{Name: "Donald"}

    // 将 Duck 实例传递给 makeItQuack 函数
    makeItQuack(myDuck)
}
