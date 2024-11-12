package main

import (
	"fmt"
)

func main() {
	var counter int
	largest := -1 << 31
	secondLargest := -1 << 31

	for counter < 10 {
		var number int
		fmt.Print("Enter number: ")
		fmt.Scan(&number)

		if number > largest {
			secondLargest = largest
			largest = number
		} else if number > secondLargest {
			secondLargest = number
		}

		counter++
	}

	fmt.Printf("The largest number is: %d\n", largest)
	fmt.Printf("The second largest number is: %d\n", secondLargest)
}
