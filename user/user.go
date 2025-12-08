package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	dateOfbirth string
	createdAt   time.Time
}
type Admin struct {
	email    string
	password string
	User
}

// attaching a function to a struct i.e method
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.dateOfbirth)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor functions
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{ //struct embedding
			firstName:   "ADMIN",
			lastName:    "ADMIN",
			dateOfbirth: "---",
			createdAt:   time.Now(),
		},
	}
}

func New(firstName, lastName, dateOfbirth string) (*User, error) {
	if firstName == "" || lastName == "" || dateOfbirth == "" {
		return nil, errors.New("firstName, lastName, dateOfBirth are required.")
	}
	return &User{
		firstName:   firstName,
		lastName:    lastName,
		dateOfbirth: dateOfbirth,
		createdAt:   time.Now(),
	}, nil
}
