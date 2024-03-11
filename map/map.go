package main

import (
	"fmt"
	"sort"
)

func main() {
    // 创建一个空的 map
    m := make(map[string]int)

    // 向 map 中添加键值对
    m["apple"] = 10
    m["banana"] = 20
    m["pear"] = 18
    m["orange"] = 15

    // 获取 map 中的值
    fmt.Println("Value of apple:", m["apple"])

    // 检查 map 中是否存在某个键
    _, ok := m["banana"]
    fmt.Println("Is banana in the map?", ok)

    // 删除 map 中的键值对
    delete(m, "orange")

    // 遍历 map 中的所有键值对
    fmt.Println("All key-value pairs in the map:")
    for key, value := range m {
        fmt.Println(key, ":", value)
    }

    // Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序
    // 如果要按顺序遍历key/value对，必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序
    var names []string
    for name := range m {
        names = append(names, name)
    }
    sort.Strings(names)

    fmt.Println("Afer sorting all key-value pairs in the map:")
    for _, name := range names {
        fmt.Printf("%s\t%d\n", name, m[name])
    }

    // map类型的零值是nil，也就是没有引用任何映射Map
    var ages map[string]int
    fmt.Println(ages == nil)  // "true"
    fmt.Println(len(ages) == 0) // "true"

    // 注意事项：
    // 1. map 是无序的，遍历结果的顺序可能不固定。
    // 2. 使用 delete 函数删除 map 中的键值对。
    // 3. 当尝试获取 map 中不存在的键时，会返回零值，但同时也会返回一个布尔值指示键是否存在。
    // 4. 对于引用类型的值，比如切片、映射等，直接赋值给其他变量时，两者共享底层数据，修改一个会影响另一个。
}

/*
$ go run .
Value of apple: 10
Is banana in the map? true
All key-value pairs in the map:
banana : 20
pear : 18
apple : 10
Afer sorting all key-value pairs in the map:
apple   10
banana  20
pear    18
true
true

1. 首先，使用 `make` 函数创建了一个空的映射 `m`，键为字符串类型，值为整数类型
2. 使用键作为索引，将键值对添加到映射中，其中键是水果的名称，值是对应水果的数量
3. 通过 `fmt.Println` 函数打印了映射中某个键的值，并检查了映射中是否存在某个键
4. 使用 `delete` 函数删除了映射中的一个键值对
5. 使用 `for range` 循环遍历了整个映射，并打印了每个键值对
6. 如果要按顺序遍历key/value对，必须显式地对key进行排序
7. map类型的零值是nil，也就是没有引用任何映射Map，如果使用 `make` 赋值后便不再是nil
*/
