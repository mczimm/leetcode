package main

import "fmt"

func main() {
	fmt.Println(addDigits(38))  //2
	fmt.Println(addDigits(0))   //0
	fmt.Println(addDigits2(38)) //2
	fmt.Println(addDigits2(0))  //0
}

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	for {
		div1 := num % 10
		div2 := num / 10
		num = div1 + div2
		if num < 10 {
			return num
		}
	}
}

func addDigits2(num int) int {
	return num % 9
}
