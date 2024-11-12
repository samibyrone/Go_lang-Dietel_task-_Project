//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var n int
//	fmt.Print("Enter a nonnegative integer: ")
//	fmt.Scan(&n)
//	fmt.Printf("Factorial of %d is %d\n", n, factorial(n))
//}
//
//func factorial(n int) int {
//	if n == 0 {
//		return 1
//	}
//	result := 1
//	for i := 1; i <= n; i++ {
//		result *= i
//	}
//	return result
//}
//
//
//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	var terms int
//	fmt.Print("Enter the number of terms to calculate e: ")
//	fmt.Scan(&terms)
//	fmt.Printf("Estimated value of e: %.10f\n", estimateE(terms))
//}
//
//func estimateE(terms int) float64 {
//	e := 1.0
//	for i := 1; i < terms; i++ {
//		e += 1.0 / float64(factorial(i))
//	}
//	return e
//}
//
//func factorial(n int) int {
//	if n == 0 {
//		return 1
//	}
//	result := 1
//	for i := 1; i <= n; i++ {
//		result *= i
//	}
//	return result
//}
//
//
//
//package main
//
//import (
//"fmt"
//"math"
//)
//
//func main() {
//	var x float64
//	var terms int
//	fmt.Print("Enter the value of x: ")
//	fmt.Scan(&x)
//	fmt.Print("Enter the number of terms to calculate e^x: ")
//	fmt.Scan(&terms)
//	fmt.Printf("Estimated value of e^%.2f: %.10f\n", x, calculateEx(x, terms))
//}
//
//func calculateEx(x float64, terms int) float64 {
//	ex := 1.0
//	for i := 1; i < terms; i++ {
//		ex += math.Pow(x, float64(i)) / float64(factorial(i))
//	}
//	return ex
//}
//
//func factorial(n int) int {
//	if n == 0 {
//		return 1
//	}
//	result := 1
//	for i := 1; i <= n; i++ {
//		result *= i
//	}
//	return result
//}
