package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	input, _ := reader.ReadString('\n')
	var message string = greeter(input)
	fmt.Println(message)

}

/**
 * This method takes an argument of type string and returns a greeting message.
 * @param {string} name
 * @return {string}
 */
func greeter(name string) string {
	return "Hello " + name
}
