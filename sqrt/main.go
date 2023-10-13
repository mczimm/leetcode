/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(4)) //2
	fmt.Println(mySqrt(8)) //2
	fmt.Println(mySqrt(0)) //2
}

func mySqrt(x int) int {
	t := findRoundedSquare(float64(x))
	return t
}

func findRoundedSquare(number float64) int {
	var sr float64
	sr = number / 2 //0
	var temp float64
	for {
		temp = sr                         //0
		sr = (temp + (number / temp)) / 2 //0
		if sr != sr {
			sr = 0
		}
		if (temp - sr) == 0 {
			break
		}
	}
	return int(sr)
}
