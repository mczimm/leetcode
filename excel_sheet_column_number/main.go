package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("A")) //1
}

// From the solutions
func titleToNumber(columnTitle string) int {
	sum := 0
	for _, s := range columnTitle {
		sum *= 26
		sum += int(s) - 64
	}
	return sum
}
