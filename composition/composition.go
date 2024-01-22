package main

import "fmt"

// Animal 类型
type Animal struct {
    Name string
}

// Speak 方法
func (a *Animal) Speak() {
    fmt.Println("Animal speaks")
}

// Dog 类型嵌入 Animal 类型
type Dog struct {
    Animal
    Breed string
}

// Speak 方法的覆盖
func (d *Dog) Speak() {
    fmt.Println("Dog barks")
}

func main() {
    // 创建一个 Dog 实例
    myDog := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }

    // 调用嵌入类型 Animal 的 Speak 方法
    myDog.Animal.Speak()

    // 调用覆盖后的 Dog 的 Speak 方法
    myDog.Speak()
}
