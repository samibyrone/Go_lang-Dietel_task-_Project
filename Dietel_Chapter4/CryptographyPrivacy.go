//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var number int
//	fmt.Print("Enter a four-digit integer to encrypt: ")
//	fmt.Scan(&number)
//
//	digits := make([]int, 4)
//	for i := 3; i >= 0; i-- {
//		digits[i] = (number%10 + 7) % 10
//		number /= 10
//	}
//
//	// Swap first and third, second and fourth digits
//	//digits[0], digits[2] = digits[2], digits[0]
//	//digits[1], digits[3] = digits[3], digits[1]
//
//	fmt.Print("Encrypted number: ")
//	for _, digit := range digits {
//		fmt.Print(digit)
//	}
//	fmt.Println()
//}
//
//
//
//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	var number int
//	fmt.Print("Enter a four-digit encrypted integer to decrypt: ")
//	fmt.Scan(&number)
//
//	digits := make([]int, 4)
//	for i := 3; i >= 0; i-- {
//		digits[i] = number % 10
//		number /= 10
//	}
//
//	// Swap first and third, second and fourth digits
//	//digits[0], digits[2] = digits[2], digits[0]
//	//digits[1], digits[3] = digits[3], digits[1]
//
//	for i := 0; i < 4; i++ {
//		digits[i] = (digits[i] + 10 - 7) % 10
//	}
//
//	fmt.Print("Decrypted number: ")
//	for _, digit := range digits {
//		fmt.Print(digit)
//	}
//	fmt.Println()
//}