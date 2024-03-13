package main

import (
	"fmt"
	"reflect"
)

type Rect struct {
	width, length float64
}

func (rect Rect) area() float64 {
	return rect.width * rect.length
}

// 在Golang中结构体和数据库表的映射关系的建立是通过struct Tag来实现的
type User struct {
	Name   string `json:"name"` //这引号里面的就是tag
	Passwd int    `json:"passwd"`
}

func testTag() {
	user := &User{"keke", 123456}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag.Get("json")) //将tag输出出来
	}
}

func testAddr() {
	b := 200
	a := &b
	fmt.Println("the address of b:", a)
	fmt.Println("the value of b:", *a)

	var p *int //p的类型是[int型的指针]
	p = &b     //p的值为 [b的地址]
	fmt.Printf("b=%d,p=%x,*p=%d,*p=%d \n", b, p, p, *p)

	*p = 5 // *p的值为[[b的地址]的指针] (其实就是b)，这行代码也就等价于b= 5
	fmt.Printf("b=%d,p=%d,*p=%d\n", b, p, *p)

    k1 := *&*&*&*&a
    t1 := &*&*&*&*a
    fmt.Println(k1, t1)

    k2 := *&*&*&*&b
    //t2 := &*&*&*&*b //invalid operation: cannot indirect b (variable of type int)
    fmt.Println(k2)
}

func main() {
	var rect = Rect{100, 200}
	fmt.Println("Width:", rect.width, "Length:", rect.length,
		"Area:", rect.area())

	fmt.Println("======")
	testTag()
	fmt.Println("======")
	testAddr()
}

/*
$ go run .
Width: 100 Length: 200 Area: 20000
name
passwd

Go函数的参数传递方式是值传递，这句话对结构体也是适用的，想要改变结构参数数据，可以传递指针

结构体“内部函数”定义方法
虽然main函数中的rect变量可以直接调用函数area()来获取矩形面积，但是area()函数确实没有定义在Rect结构体内部
Go使用组合函数的方式,来为结构体定义结构体方法。

我们仔细看一下上面的area()函数定义
首先是关键字func表示这是一个函数，第二个参数是结构体类型和实例变量，第三个是函数名称，第四个是函数返回值
这里我们可以看出area()函数和普通函数定义的区别就在于area()函数多了一个结构体类型限定
这样一来Go就知道了这是一个为结构体定义的方法

这里需要注意一点就是定义在结构体上面的函数(function)一般叫做方法(method)



在Golang中很多时候都是需要用指针结合结构体开发的，通常指针是存储一个变量的内存地址的变量

指针不参与计算但是可以用来获取地址的，例如变量a的内存地址为&a，这里的&就是获取a的地址
如果一个指针，它的值是在别的地方的地址，而且我们想要获取这个地址的值，可以使用*符号
*符号是为取值符。例如上面的&a是一个地址，那么这个地址里存储的值为*&a

& 和 *符号的区别
这里要注意 & 和 *符号的区别，& 运算符用来获取指针地址，而* 运算符是用来获取地址里存储的值
此外指针的值和指针的地址是不同的概念，指针的值: 指的是一个地址，是别的内存地址。指针的地址: 指的是存储指针内存块的地址

＆和 * 是可以互相抵消的，但要注意，*&可以抵消掉，但&*是不一定抵消的, a和*&a是一样的，都是a的值，因为* &互相抵消掉了，同理a和*&*&*&*&a是一样的 (因为4个*&互相抵消掉了)
解释下&*是不一定抵消的，对于指针变量来说，在前面添加&*是可以抵消的，但是对于普通值变量这样做会编译错误


Golang中string，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针
若函数需改变slice的长度，则仍需要取地址传递指针

如果要访问指针 p 指向的结构体中某个元素 x，不需要显式地使用 * 运算，可以直接 p.x
*/
