package main

import "fmt"

func main() {
	// Examples of looping
	// for loop
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// while loop
	i := 0
	for i < len(days) {
		fmt.Println(days[i])
		i++
	}

	// do-while loop
	i = 0
	for {
		if i >= len(days) {
			break
		}
		fmt.Println(days[i])
		i++
	}

	// for range loop
	for i, day := range days {
		fmt.Println(i+1, day)
	}

	// for range loop with index and value
	for i, day := range days {
		fmt.Println(i+1, day)
	}

	// goto loop
	i = 0
	for {
		if i >= len(days) {
			break
		}
		if days[i] == "Friday" {
			goto Friday
		}
		i++
	}

Friday:
	fmt.Println("Today is Friday")
}
