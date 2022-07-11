package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	// Comma ok syntax || Comma error syntax
	input, _ := reader.ReadString('\n') // In here we ecpect a user input
	fmt.Println("Your name is:", input)
	fmt.Printf("The type of input is: %T \n", input)
}
