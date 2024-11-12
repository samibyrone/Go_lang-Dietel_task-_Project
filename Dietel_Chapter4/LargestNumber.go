package main

import (
	"fmt"
)

func main() {
	var counter int
	largest := -1 << 31

	for counter < 10 {
		var number int
		fmt.Print("Enter number: ")
		fmt.Scan(&number)

		if number > largest {
			largest = number
		}

		counter++
	}

	fmt.Printf("The largest number is: %d\n", largest)
}
