package models

type User struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
}

func NewUser() User {
	return User{
		Id:        1,
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
		Id:        1,
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func NewUserById(
	id string) (*User, error) {

	return &User{
		Id:        1,
		FirstName: "Bill",
		LastName:  "Mourtzis",
		Age:       29,
	}, nil
}
