package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int

	for {
		fmt.Print("Enter a five-digit integer: ")
		fmt.Scan(&number)

		if len(strconv.Itoa(number)) != 5 {
			fmt.Println("Error: The number must be exactly five digits long.")
		} else {
			break
		}
	}

	if isPalindrome(number) {
		fmt.Printf("%d is a palindrome.\n", number)
	} else {
		fmt.Printf("%d is not a palindrome.\n", number)
	}
}

func isPalindrome(number int) bool {
	str := strconv.Itoa(number)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
