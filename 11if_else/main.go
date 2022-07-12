package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrorInput int = 0

func main() {
	if ErrorInput > 3 {
		fmt.Println("Too many wrong inputs. Exiting...")
		return
	}
	fmt.Println("Welcome to the Youth club!")

	var ageRestriction int = 18
	fmt.Print("Enter your age: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	age, _ := strconv.Atoi(input)

	if age < ageRestriction {
		fmt.Println("You are not allowed to enter. You are too young!")
		ErrorInput++
		main()
	} else {
		fmt.Println("Welcome to the youth club!")
	}
}
