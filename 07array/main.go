package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array playground!")

	// Declare an array of strings with a length of 3
	// Assign the values "apple", "banana", orange" , "grape" and "pear" to the array
	fruitList := [5]string{"apple", "banana", "orange", "grape", "pear"}
	// Print the array
	fmt.Println("Fruit list:", fruitList)
	// Print the length of the array
	fmt.Println("Length of the array:", len(fruitList))
	// Print the element at index 1
	fmt.Println("Element at index 1:", fruitList[1])

	// Other ways to declare an array

	// Declare an array of integers with a length of 5
	var integerList [5]int
	// Assign the values 1, 2, 3, 4, and 5 to the array
	integerList[0] = 1
	integerList[1] = 2
	integerList[2] = 3
	integerList[3] = 4
	integerList[4] = 5
	// Print the array
	fmt.Println("Integer list:", integerList)
	// Print the length of the array
	fmt.Println("Length of the array:", len(integerList))
	// Print the element at index 1
	fmt.Println("Element at index 1:", integerList[1])
}
