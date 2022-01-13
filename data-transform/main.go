package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/jinzhu/copier"
	"github.com/qa-tools-family/grpc-demo/data-transform/data"
	"github.com/qa-tools-family/grpc-demo/data-transform/pb"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)


func pb2jsonBase()  {

	var timeProto *timestamppb.Timestamp
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

func data2pb() {
	address := data.AddressBook{People: []data.Person{
		{
			Name:   "missshi",
			Id:     0,
			Email:  "wangzhe0912@tju.edu.cn",
			Phones: nil,
			LastUpdated: time.Now(),
		},
		{
			Name:   "missshi",
			Id:     20,
			Email:  "",
			Phones: []data.PhoneNumber{
				{
					Number: "152xxxx0001",
					Type: 1,
				},
			},
		},
	}}
	logger.Info(address)
	var pbAddress pb.AddressBook
	//_ = copier.Copy(&pbAddress, &address)
	_ = copier.CopyWithOption(&pbAddress, &address, copier.Option{
		IgnoreEmpty: false,
		DeepCopy:    false,
	})
	logger.Info(&pbAddress)

	var newAddress data.AddressBook
	_ = copier.CopyWithOption(&newAddress, &pbAddress, copier.Option{
		IgnoreEmpty: false,
		DeepCopy:    false,
	})
	logger.Info(&newAddress)
}

func main() {
	//pb2jsonBase()
	//pb2jsonPb()
	data2pb()
}