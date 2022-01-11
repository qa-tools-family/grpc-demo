package main

import (
	pb "github.com/qa-tools-family/grpc-demo/protobuf/tutorialpb"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	fname := "data.pb"
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	book := &pb.AddressBook{}
	book.People = []*pb.Person{&p}
	out, err := proto.Marshal(book)
	// save data
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
	// load data
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book = &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	} else {
		log.Print(book)
	}
}
