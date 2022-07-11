package main

import "fmt"

func main() {
	fmt.Println("Welcome to the playground of slices!")

	// Declare a slice of strings
	// A slice is an abstraction of array under the hood.
	// Arrays are fixed-size, slices are variable-size.
	// array fruitList = [5]string{"apple", "banana", "orange", "grape", "pear"}
	// slice fruitList = []string{"apple", "banana", "orange", "grape", "pear"}
	var fruitList = []string{"apple", "banana", "orange", "grape", "strawberry"}
	fmt.Println("fruit list:", fruitList)
	fmt.Printf("Type of fruitList: %T\n", fruitList)

	// Now add some more fruits to the slice
	fruitList = append(fruitList, "kiwi", "pear", "mango")
	fmt.Println("New fruit list:", fruitList)

	// Now remove the 6th fruit kiwi from the slice
	fruitList = append(fruitList[5:6])
	fmt.Println("Sliced fruit list:", fruitList)
}
