package data

import (
	"encoding/json"
	"time"
)

type PhoneType int32

const (
	MOBILE PhoneType = 0
	HOME   PhoneType = 1
	WORK   PhoneType = 2
)

type PhoneNumber struct {
	Number string
	Type PhoneType
}

type Person struct {
	Name string
	Id int
	Email string
	Phones []PhoneNumber
	LastUpdated time.Time
}

type AddressBook struct {
	People []Person
}

func (a AddressBook) String() string {
	message, err := json.Marshal(a.People)
	if err != nil {
		return "err"
	} else {
		return string(message)
	}
}