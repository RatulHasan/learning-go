package main

import "fmt"

func main() {
	var username string = "Ratul Hasan"
	var isLoggedIn bool = false
	age := 30
	var height float64 = 5.8
	var weight float64 = 60.0
	var isMarried bool = true

	fmt.Println("Hello, " + username)
	fmt.Printf("The type of username is: %T \n", username)
	fmt.Println("The user is logged in:", isLoggedIn)
	fmt.Printf("The type of isLoggedIn is: %T \n", isLoggedIn)
	fmt.Println("The user's age is:", age)
	fmt.Printf("The type of age is: %T \n", age)
	fmt.Println("The user's height is:", height)
	fmt.Printf("The type of height is: %T \n", height)
	fmt.Println("The user's weight is:", weight)
	fmt.Printf("The type of weight is: %T \n", weight)
	fmt.Println("The user is married:", isMarried)
	fmt.Printf("The type of isMarried is: %T \n", isMarried)
}
