package data

import "time"

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
	Phones []*PhoneNumber
	last_updated *time.Time
}

type AddressBook struct {
	People []*Person
}
