package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the rating app!")
	fmt.Println("Please enter your rating between 1 and 5:")

	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')

	fmt.Println("Your rating is:", rating)
	// The type of input is: string
	fmt.Printf("Thank you for your input! The type of input is: %T \n", rating)

	// Convert the string to an int
	intRating := strings.TrimSpace(rating)
	// Convert the string to float64
	ratingInt, error := strconv.ParseFloat(intRating, 64)

	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Your rating is:", ratingInt/2)
	}
}
