package main

import "fmt"

func main() {
	phoneNumber := 1234

	fmt.Println("Phone number:", phoneNumber)

	// pointer to phoneNumber
	pointer := &phoneNumber
	fmt.Println("pointer:", pointer)
	fmt.Println("pointer value:", *pointer)

	// change the value of phoneNumber using pointer
	*pointer = 4321
	fmt.Println("Phone number:", phoneNumber)
}
