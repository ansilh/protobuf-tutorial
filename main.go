package main

import (
	"github.com/ansilh/protobuf-tutorial/pb"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {

	ansil := &pb.Person{
		Name: "Ansil",
		Age:  34,
	}
	fname := "person.pbuf"

	data, err := proto.Marshal(ansil)

	if err != nil {
		log.Fatalln("Marshalling error:", err)
	}

	if err := ioutil.WriteFile(fname, data, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
