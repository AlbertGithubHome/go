package main

import (
	"fmt"
	"log"

	msg "lib_protobuf/msgproto"

	"google.golang.org/protobuf/proto"
)

func main() {
	p := &msg.Person{
		Name:    "Alice",
		Age:     30,
		Hobbies: []string{"reading", "running"},
	}

	// 序列化
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatalf("Failed to encode person: %v", err)
	}

	// 反序列化
	newP := &msg.Person{}
	err = proto.Unmarshal(data, newP)
	if err != nil {
		log.Fatalf("Failed to decode person: %v", err)
	}

	fmt.Println("Original Person:", p)
	fmt.Println("Decoded Person:", newP)
}

/*
$ go run .
Original Person: name:"Alice"  age:30  hobbies:"reading"  hobbies:"running"
Decoded Person: name:"Alice"  age:30  hobbies:"reading"  hobbies:"running"

在这个小工程中，srcproto/examplremsg.proto相当于一个配置文件，可以不包含在这个工程中，
而 msgproto/examplremsg.pb.go 是.proto文件的导出文件，需要包含在工程中来避免编译错误，
如果在每次编译之前可以重新导出，这些导出文件也不必包含在工程中
*/
