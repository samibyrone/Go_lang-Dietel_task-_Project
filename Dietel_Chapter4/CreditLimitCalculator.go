package main

import (
	"fmt"
)

func main() {
	for {
		var accountNumber, beginningBalance, totalCharges, totalCredit, creditLimit int

		fmt.Print("Enter Account Number(0 to quit): ")
		fmt.Scan(&accountNumber)
		if accountNumber == 0 {
			break
		}

		fmt.Print("Enter Beginning Balance for the month: ")
		fmt.Scan(&beginningBalance)

		fmt.Print("Enter Total items charged for the month: ")
		fmt.Scan(&totalCharges)

		fmt.Print("Enter Total Credit applied for the month: ")
		fmt.Scan(&totalCredit)

		fmt.Print("Enter Credit card limit: ")
		fmt.Scan(&creditLimit)

		newbalance := beginningBalance + totalCharges - totalCredit

		fmt.Printf("New balance for the month: %d\n", newbalance)

		if newbalance > creditLimit {
			fmt.Println("Your balance exceeded your credit limit")
		} else {
			fmt.Println("Your balance does not exceeded your credit limit")
		}
	}
}
