package main

import (
	"fmt"
)

func main() {
	var angleBaseLength int

	for {
		fmt.Print("Enter the length of the base of the triangle (1-10): ")
		fmt.Scan(&angleBaseLength)

		if angleBaseLength >= 1 && angleBaseLength <= 10 {
			break
		} else {
			fmt.Println("Error: The base length must be between 1 and 10.")
		}
	}

	for i := 1; i <= angleBaseLength; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
