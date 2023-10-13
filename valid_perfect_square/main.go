package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16)) //true
	fmt.Println(isPerfectSquare(14)) //false
	fmt.Println(isPerfectSquare(1))  //true
}

func isPerfectSquare(num int) bool {
	left := 0
	right := num

	for left <= right {
		mid := (left + right) / 2
		if mid*mid > num {
			right = mid - 1
		} else if mid*mid < num {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}
