package main

import "fmt"

func main() {
	//fmt.Println(minPartitions("32"))    //3
	fmt.Println(minPartitions("82734")) //8
}

// From the solutions: https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/solutions/3441947/golang/
func minPartitions(n string) int {
	var maxDigit uint8

	for i := 0; i < len(n); i++ {
		currDigit := n[i] - '0' //int32 - uint8, returns uint8 which means rune
		if currDigit > maxDigit {
			maxDigit = currDigit
		}
	}
	return int(maxDigit)
}
