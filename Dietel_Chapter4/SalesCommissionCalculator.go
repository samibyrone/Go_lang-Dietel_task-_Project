package main

import (
	"fmt"
)

func main() {
	var numberOfItems int
	var totalSales float64
	baseSalary := 200.0
	commissionRate := 0.09

	fmt.Print("Enter the number of items sold: ")
	fmt.Scan(&numberOfItems)

	for i := 1; i <= numberOfItems; i++ {
		var itemValue float64
		fmt.Printf("Enter the value of item %d: ", i)
		fmt.Scan(&itemValue)
		totalSales += itemValue
	}

	earnings := baseSalary + (commissionRate * totalSales)
	fmt.Printf("Total earnings: $%.2f\n", earnings)
}
