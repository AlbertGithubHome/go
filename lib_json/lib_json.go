package main

import (
	"encoding/json"
	"fmt"
)

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Channel  int    `json:"channel"`
}

func main() {
	// 创建 LoginRequest 实例
	request := LoginRequest{
		Account:  "alberts@example.com",
		Password: "password123",
		Channel:  1,
	}

	// 序列化为 JSON 字符串
	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印 JSON 字符串
	fmt.Println("Serialized JSON:", string(jsonData))
}

/*
$ go run .
Serialized JSON: {"account":"alberts@example.com","password":"password123","channel":1}

在上面的示例中，LoginRequest 结构体的字段上使用了 json 标签，指定了 JSON 字段的名称。
json.Marshal 函数会根据这些标签将结构体转换为 JSON 字符串。

encoding/json 是 Go 语言的标准库之一，用于处理 JSON 数据的编码（序列化）和解码（反序列化）。
因此，不需要额外安装任何包，可以直接在你的 Go 项目中使用它。


除了 `json` 标签，Go 中的结构体字段还可以使用其他标签，这些标签的作用取决于具体的使用场景。
以下是一些常见的标签及其用途：

1. **`xml` 标签：** 用于指定 XML 编码和解码时的字段名称和属性。例如：

    ```go
    type Person struct {
        Name  string `xml:"name"`
        Age   int    `xml:"age"`
        City  string `xml:"city"`
    }
    ```

2. **`db` 标签：** 用于指定数据库字段的名称和属性。在使用数据库 ORM（Object-Relational Mapping）库时，
这些标签可能会用于与数据库表的映射。例如：

    ```go
    type User struct {
        ID    int    `db:"id"`
        Name  string `db:"name"`
        Email string `db:"email"`
    }
    ```

3. **`form` 标签：** 用于指定 HTML 表单中字段的属性。在处理 web 表单时，
这些标签可能被用于生成 HTML 表单或验证用户输入。例如：

    ```go
    type LoginForm struct {
        Username string `form:"username"`
        Password string `form:"password"`
    }
    ```

4. **`yaml` 标签：** 用于指定 YAML 编码和解码时的字段名称。与 `json` 标签类似，
但用于 YAML 格式。例如：

    ```go
    type Configuration struct {
        APIKey string `yaml:"api_key"`
        Port   int    `yaml:"port"`
    }
    ```

这些标签的使用是可选的，取决于具体的需求和场景。标签可以帮助在不同领域的数据表示之间建立映射关系，
使得编码和解码更加灵活。标签的名称和含义通常由使用它们的库或工具来定义。
*/
