package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange2([]int{1, 2, 5}, 11))            // 3
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

func coinChange2(coins []int, amount int) int {
	// Initialize the DP array with a large value (infinity).
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0 // Base case: no coins are needed to make amount 0.

	// Fill the DP array
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	// If dp[amount] is still the large value, return -1 (impossible case).
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// Helper function to find the minimum of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
