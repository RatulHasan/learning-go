package main

import "fmt"

func main() {
	// Defer means first came last out.
	defer fmt.Println("One more line")
	fmt.Println("And another line")
	fmt.Println("Hello, world!")

	myDefer()

	// Another example of defer
	defer fmt.Println("Start")
	fmt.Println("End")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
