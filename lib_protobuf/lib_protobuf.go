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
