package main

import (
	"fmt"
)

func main() {
	var input int

	for input != 1 && input != 2 {
		fmt.Print("Enter 1 or 2: ")
		fmt.Scan(&input)

		if input != 1 && input != 2 {
			fmt.Println("Invalid input. Please enter 1 or 2.")
		}
	}

	fmt.Printf("You entered a valid input: %d\n", input)
}
