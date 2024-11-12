package main

import (
	"fmt"
)

func main() {
	var firstNumber, secondNumber int

	fmt.Print("Enter the first number: ")
	fmt.Scan(&firstNumber)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&secondNumber)

	result := compareNumbers(firstNumber, secondNumber)
	fmt.Printf("Comparison result: %d\n", result)
}

func compareNumbers(firstNumber, secondNumber int) int {
	if firstNumber == secondNumber {
		return 0
	} else if firstNumber > secondNumber {
		return 1
	} else {
		return -1
	}
}
