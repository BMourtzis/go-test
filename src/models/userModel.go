package models

import (
	"errors"

	"github.com/rs/xid"
)

type User struct {
	Uid       xid.ID
	FirstName string
	LastName  string
	Age       int
}

func NewUser() User {
	return User{
		Uid:       xid.New(),
		FirstName: "Bill",
		LastName:  "Mourtzis",
		Age:       29,
	}
}

func NewUserWithParameters(
	firstName string,
	lastName string,
	age int) User {
	return User{
		Uid:       xid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func NewUserById(
	id string) (*User, error) {
	uid, err := xid.FromString(id)

	if err != nil {
		return nil, errors.New("incorrect id")
	}

	return &User{
		Uid:       uid,
		FirstName: "Bill",
		LastName:  "Mourtzis",
		Age:       29,
	}, nil
}
