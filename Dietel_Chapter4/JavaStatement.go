package main

import (
	"fmt"
)

func main() {
	x := 5
	x += x - 5
	x++
	fmt.Printf("The value of x is: %d\n", x)
}
