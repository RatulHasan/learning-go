package main

import (
	"fmt"
)

func main() {
	newUser := User{
		FullName: "John Doe",
		UserName: "jdoe",
		Email:    "ratuljh@gmail.com",
		Password: nil,
		Status:   true,
	}
	newUser.CreatePassword("password")
	fmt.Printf("%+v\n", newUser)
}

// User type struct
type User struct {
	FullName string
	UserName string
	Email    string
	Password []byte
	Status   bool
}

// CreatePassword with md5 hash
// Make sure to use pointer to the user struct
func (u *User) CreatePassword(password string) User {
	u.Password = []byte(password)
	return *u
}
