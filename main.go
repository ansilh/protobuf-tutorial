package main

import (
	"fmt"
	"github.com/ansilh/protobuf-tutorial/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	fmt.Println("Hello World")

	ansil := &pb.Person{
		Name: "Ansil",
		Age:  34,
	}

	data, err := proto.Marshal(ansil)
	if err != nil {
		log.Fatal("Marshalling error: %v", err)
	}

	fmt.Println(data)
}
