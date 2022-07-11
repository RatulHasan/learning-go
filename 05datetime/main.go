package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the date time app!")

	presentTime := time.Now()
	// Formatting time to day month year.
	// Remember the format you have to use is "01-02-2006" Monday
	fmt.Println("Todays date is: ", presentTime.Format("01-02-2006 Monday"))
}
