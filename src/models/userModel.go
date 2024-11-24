package models

import "github.com/rs/xid"

type User struct {
	Uid      xid.ID
	FistName string
	LastName string
	Age      int
}

func NewUser() User {
	return User{
		Uid:      xid.New(),
		FistName: "Bill",
		LastName: "Mourtzis",
		Age:      29,
	}
}

func NewUserWithParameters(
	firstName string,
	lastName string,
	age int) User {
	return User{
		Uid:      xid.New(),
		FistName: firstName,
		LastName: lastName,
		Age:      age,
	}
}
