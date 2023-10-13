package main

import "fmt"

func main() {
	fmt.Println(guessNumber(10)) //6
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 0, n
	var mid int

	for left <= right {
		mid = (left + right) / 2
		if guess(mid) == -1 {
			right = mid - 1
		} else if guess(mid) == 1 {
			left = mid + 1
		} else {
			break
		}
	}
	return mid
}

func guess(num int) int {
	if num > 6 {
		return -1
	} else if num < 6 {
		return 1
	} else {
		return 0
	}
}
