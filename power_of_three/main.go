package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfThree(27)) //true
	fmt.Println(isPowerOfThree(45)) //false
	fmt.Println()
	fmt.Println(isPowerOfThree2(27)) //true
	fmt.Println(isPowerOfThree2(45)) //false
}

// Recursion
func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n < 1 || n%3 != 0 {
		return false
	}
	return isPowerOfThree(n / 3)
}

// Loop
func isPowerOfThree2(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
