package main

import (
	"fmt"
)

func main() {
	var totalMiles, totalGallons int

	for {
		var miles, gallons int
		fmt.Print("Enter the Miles Driven (-1 to quit): ")
		fmt.Scan(&miles)
		if miles == -1 {
			break
		}

		fmt.Print("Enter the Gallons used for Driving: ")
		fmt.Scan(&gallons)

		totalMiles += miles
		totalGallons += gallons

		tripMilesPerGallon := float64(miles) / float64(gallons)
		totalMilesPerGallon := float64(totalMiles) / float64(totalGallons)

		fmt.Printf("Miles per Gallon used for this trip (%.2f)\n", tripMilesPerGallon)
		fmt.Printf("Combined Miles per Gallon (%.2f)\n", totalMilesPerGallon)
	}
}
