package main

import (
	"fmt"
)

func main() {
	currentPopulation := 8.2e9
	growthRate := 0.0091
	years := 75

	fmt.Printf("%-10s %-20s %-20s\n", "Year", "Population", "Increase")
	for year := 1; year <= years; year++ {
		increase := currentPopulation * growthRate
		currentPopulation += increase
		fmt.Printf("%-10d %-20.2f %-20.2f\n", year, currentPopulation, increase)
	}

	initialPopulation := 8.2e9
	currentPopulation = initialPopulation
	doublingYear := 0
	for year := 1; year <= years; year++ {
		increase := currentPopulation * growthRate
		currentPopulation += increase
		if currentPopulation >= 2*initialPopulation && doublingYear == 0 {
			doublingYear = year
		}
	}
	fmt.Println("The population would double in year:", doublingYear)
}
