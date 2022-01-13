package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/qa-tools-family/grpc-demo/data-transform/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)


func pb2jsonBase()  {

	var timeProto *timestamp.Timestamp
	timeProto = timestamppb.Now()

	person := pb.Person{
		Name:        "missshi",
		Id:          5,
		Email:       "wangzhe0912@tju.edu.cn",
		Phones:      []*pb.PhoneNumber{
			{
				Number: "152xxxx1111",
				Type: pb.PhoneType_HOME,
			},
		},
		LastUpdated: timeProto,
	}
	address := pb.AddressBook{People: []*pb.Person{&person}}
	jsonData, err := json.Marshal(&address)
	if err != nil {
		fmt.Println("json dumps address failed: ", err)
	}
	fmt.Println(string(jsonData))

	var newAddress *pb.AddressBook
	err = json.Unmarshal(jsonData, &newAddress)
	if err != nil {
		fmt.Println("Unmarshal failed: ", err)
	} else {
		fmt.Println(newAddress)
	}
}

func pb2jsonPb() {

	person := pb.Person{
		Name:        "missshi",
		Id:          5,
		Email:       "wangzhe0912@tju.edu.cn",
		Phones:      []*pb.PhoneNumber{
			{
				Number: "152xxxx1111",
				Type: pb.PhoneType_HOME,
			},
		},
		LastUpdated: nil,
	}
	address := pb.AddressBook{People: []*pb.Person{&person}}
	marshal := jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "    ",
		AnyResolver:  nil,
	}
	jsonData, err := marshal.MarshalToString(&address)
	if err != nil {
		fmt.Println("json dumps address failed: ", err)
	}
	fmt.Println(jsonData)

	var newAddress pb.AddressBook
	unmarshal := jsonpb.Unmarshaler{
		AllowUnknownFields: true,
		AnyResolver:        nil,
	}
	err = unmarshal.Unmarshal(bytes.NewReader([]byte(jsonData)), &newAddress)
	if err != nil {
		fmt.Println("json load address failed: ", err)
	} else {
		fmt.Println(&newAddress)
	}
}

func main() {
	pb2jsonBase()
	pb2jsonPb()
}