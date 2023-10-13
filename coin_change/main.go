package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))             // 3
	fmt.Println(coinChange([]int{2, 5, 10, 1}, 27))         // 4
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249)) // 20
}

func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}

	result := make([]int, amount+1)
	for i := range result {
		result[i] = math.MaxInt32
	}
	result[0] = 0

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			result[j] = Min(result[j], result[j-coins[i]]+1)
		}
	}

	if result[amount] == math.MaxInt32 {
		return -1
	}
	return result[amount]
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
