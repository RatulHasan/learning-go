package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var totalDiceNumber int = 0

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press ENTER to roll the dice!")
	_, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	totalDiceNumber = totalDiceNumber + diceNumber
	fmt.Println("You rolled a", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("You can go 1 step forward!")
	case 2:
		fmt.Println("You can go 2 steps forward!")
	case 3:
		fmt.Println("You can go 3 steps forward!")
	case 4:
		fmt.Println("You can go 4 steps forward!")
	case 5:
		fmt.Println("You can go 5 steps forward!")
	case 6:
		fmt.Println("You can go 6 steps forward! And you can roll the dice again!")
		fmt.Println("Total dice number:", totalDiceNumber)
		main()
	}

	fmt.Println("Total dice number:", totalDiceNumber)
	if totalDiceNumber >= 50 {
		fmt.Println("You win!")
	} else {
		main()
	}
}
