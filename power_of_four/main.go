package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPowerOfFour(16)) //true
	fmt.Println(isPowerOfFour(5))  //false
	fmt.Println(isPowerOfFour(1))  //true
}

// Follow up: Could you solve it without loops/recursion?
// My first attempt
func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}

	res := math.Log(float64(n)) / math.Log(float64(4))

	if res == math.Floor(res) {
		return true
	}
	return false
}
