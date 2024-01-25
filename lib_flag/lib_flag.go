package main

import (
	"flag"
	"fmt"
	"strconv"
)

// Custom type NewType
type NewType struct {
	Value int
}

func (m *NewType) String() string {
	return fmt.Sprintf("%d", m.Value)
}

func (m *NewType) Set(value string) error {
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	m.Value = parsedValue
	return nil
}

func main() {
	// String parameter
	var name string
	flag.StringVar(&name, "name", "default", "Specify a name")

	// Integer parameter
	var age int
	flag.IntVar(&age, "age", 0, "Specify an age")

	// Float parameter
	var salary float64
	flag.Float64Var(&salary, "salary", 0.0, "Specify a salary")

	// Boolean parameter
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose mode")

	// Custom type parameter
	var customParam NewType
	flag.Var(&customParam, "custom", "Specify a custom parameter")

	// Parse command-line arguments
	flag.Parse()

	// Print parameter values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Verbose:", verbose)
	fmt.Println("Custom Param:", customParam)
}

/*
运行结果
$ go run . -name John -age 25 -salary 50000.5 -verbose -custom 180
Name: John
Age: 25
Salary: 50000.5
Verbose: true
Custom Param: {180}
*/

/*
flag 是 Go 语言的标准库之一，用于处理命令行参数。
flag 包提供了一种简单而灵活的方式来定义和解析命令行参数，包括标志（flags）和参数。

需要先绑定变量和预设输入的关系，然后再解析命令行参数

在 Go 语言中，标准库由一系列的包组成，这些包提供了各种功能，例如输入输出、网络、并发、解析命令行参数等。
flag 包属于这些标准库之一，因此无需额外安装即可在你的 Go 代码中使用它。
*/
