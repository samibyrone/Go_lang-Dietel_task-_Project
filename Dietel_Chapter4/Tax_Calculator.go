package main

import (
	"fmt"
)

func main() {
	const LOWER_RATE = 0.15
	const UPPER_RATE = 0.20
	const THRESHOLD = 30000.0

	for i := 0; i < 3; i++ {
		var name string
		var earnings float64

		fmt.Print("Enter citizen's name: ")
		fmt.Scan(&name)

		fmt.Printf("Enter %s's earnings: ", name)
		fmt.Scan(&earnings)

		var tax float64
		if earnings <= THRESHOLD {
			tax = earnings * LOWER_RATE
		} else {
			tax = (THRESHOLD * LOWER_RATE) + ((earnings - THRESHOLD) * UPPER_RATE)
		}

		fmt.Printf("Total tax for %s: $%.2f\n", name, tax)
	}
}
