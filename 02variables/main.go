package main

import "fmt"

func main() {
	var username string = "Ratul Hasan"
	var isLoggedIn bool = false

	fmt.Println("Hello, " + username)
	fmt.Printf("The type of username is: %T \n", username)
	fmt.Println("The user is logged in:", isLoggedIn)
	fmt.Printf("The type of isLoggedIn is: %T \n", isLoggedIn)
}
