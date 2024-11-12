package main

import (
	"fmt"
)

func main() {
	var threshold, sum int

	fmt.Print("Enter the threshold value: ")
	fmt.Scan(&threshold)

	for sum < threshold {
		var number int
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)
		sum += number
		fmt.Printf("Current sum: %d\n", sum)
	}

	fmt.Println("The sum has reached or exceeded the threshold.")
}
