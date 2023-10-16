package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(hammingDistance(1, 4))
}

// From the solutions:
func hammingDistance(x int, y int) int {
	xor := x ^ y
	res := 0

	for xor != 0 {
		xor >>= 1
		res += xor % 2
	}
	return res
}

// From the solutions: https://leetcode.com/problems/hamming-distance/solutions/3922050/go-simple-solutions-1-line/
func hammingDistance2(x, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
