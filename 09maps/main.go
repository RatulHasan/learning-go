package main

import "fmt"

func main() {
	fmt.Println("Welcome to the playground of maps!")

	// Declare a map of programming languages
	// A map is a collection of key-value pairs.
	// The keys of a map must be unique.
	// The values of a map can be of any type.
	// The keys of a map are always of the same type.
	// The values of a map are always of the same type.
	// map[key type]value type
	programmingLanguages := make(map[string]string)
	programmingLanguages["Go"] = "Go"
	programmingLanguages["JS"] = "JavaScript"
	programmingLanguages["RB"] = "Ruby"
	programmingLanguages["Python"] = "Python"
	programmingLanguages["Java"] = "Java"
	programmingLanguages["C++"] = "C++"
	programmingLanguages["C"] = "C"
	fmt.Println("Programming languages:", programmingLanguages)

	// Loop through the map
	for key, value := range programmingLanguages {
		fmt.Println("Key:", key, "Value:", value)
	}

	// Accessing a value of a key in a map.
	fmt.Println("JS is a programming language:", programmingLanguages["Go"])

	// Adding a new key-value pair to the map.
	programmingLanguages["C#"] = "C#"
	fmt.Println("Programming languages:", programmingLanguages)

	// Delete a key-value pair from the map.
	delete(programmingLanguages, "RB")
	fmt.Println("Programming languages:", programmingLanguages)

	// Check if a key exists in the map.
	_, exists := programmingLanguages["JS"]
	fmt.Println("JS exists in the map:", exists)

	// Declare a map of people and their age.
	people := make(map[string]int)
	people["John"] = 30
	people["Mary"] = 25
	people["Bob"] = 40
	people["Jane"] = 35
	fmt.Println("People:", people)

	// Accessing a value of a key in a map.
	fmt.Println("John is", people["John"], "years old.")

	// Loop through the map
	for key, value := range people {
		fmt.Printf("%s is %d years old.\n", key, value)
	}

	// Delete a key-value pair from the map.
	delete(people, "John")
	fmt.Println("People:", people)

	// Check if a key exists in the map.
	_, exists = people["John"]
	fmt.Println("John exists in the map:", exists)

}
