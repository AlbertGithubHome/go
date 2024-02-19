// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(3, 5)
    expected := 8
    if result != expected {
        t.Errorf("Add(3, 5) returned %d, expected %d", result, expected)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(10, 5)
    expected := 5
    if result != expected {
        t.Errorf("Subtract(10, 5) returned %d, expected %d", result, expected)
    }
}

/*
$ go run .
package math_test is not a main package

$ go test
PASS
ok      math_test       0.001s


在这个示例中，我们定义了一个 `math` 包，包含两个函数 `Add` 和 `Subtract`，分别用于加法和减法运算。
然后，我们创建了一个名为 `math_test.go` 的测试文件，其中包含了两个测试函数 `TestAdd` 和 `TestSubtract`，
用于测试 `Add` 和 `Subtract` 函数的功能是否正确。

在每个测试函数中，我们调用相应的函数并检查返回值是否符合预期。
如果返回值与预期值不一致，我们使用 `t.Errorf` 函数输出测试失败的信息。

要运行这些测试函数，可以使用 `go test` 命令。在包含 `math_test.go` 文件的目录中
运行 `go test` 命令，测试工具会自动识别并执行测试函数，并输出测试结果。

在 Go 语言中，`testing.T` 是测试的上下文对象，用于编写和执行测试。
该对象提供了一组方法，用于报告测试失败、跳过测试等。在测试函数中，
我们通常会使用 `testing.T` 对象来检查函数的行为是否符合预期，
并通过调用 `t.Errorf`、`t.Fatalf` 等方法来报告测试失败。

在测试函数的签名中，`*testing.T` 是一个指针类型，指向测试上下文对象的实例。
通过将该对象传递给测试函数，我们可以在函数内部使用它来执行各种测试操作。
*/