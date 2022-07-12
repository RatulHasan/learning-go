package main

import "fmt"

func main() {
	fmt.Println("Welcome to the playground of structs!")

	// Let's create a user using User structs
	ratul := User{
		FullName: "Ratul Hasan",
		username: "ratul",
		Email:    "ratuljh@gmail.com",
		Password: "123456",
		Status:   true,
	}

	fmt.Printf("User: %+v\n", ratul)
}

// User is a struct that holds information about a user
// The struct will have the following fields:
// FullName: string
// Email: string
// Password: string
// Status: bool
type User struct {
	FullName string
	username string
	Email    string
	Password string
	Status   bool
}
