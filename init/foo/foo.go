package foo

import "fmt"

var x int

func init() {
    x = 18
    fmt.Println("foo package is initialized")
}

func GetX() int {
    return x
}
