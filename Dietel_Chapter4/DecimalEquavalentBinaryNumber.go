package main

import (
	"fmt"
	"math"
)

func main() {
	var binaryNumber int
	fmt.Print("Enter a binary number: ")
	fmt.Scan(&binaryNumber)

	decimalNumber := 0
	power := 0

	for binaryNumber != 0 {
		lastDigit := binaryNumber % 10
		decimalNumber += lastDigit * int(math.Pow(2, float64(power)))
		binaryNumber /= 10
		power++
	}

	fmt.Printf("Decimal equivalent: %d\n", decimalNumber)
}
