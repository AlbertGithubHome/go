package main

import (
	"fmt"
	"runtime"
	"time"
)

func switchDemo() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.\n", os)
    }
}

func main() {
    switchDemo()
    fmt.Println("=====")

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}

/*
$ go run .
Go runs on Linux.
=====
Good afternoon.

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case。
实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止。

Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。
没有条件的 switch 同 switch true 一样。这种形式能将一长串 if-then-else 写得更加清晰。
*/
