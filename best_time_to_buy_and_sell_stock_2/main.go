package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 7
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))    // 4
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // 0
}

// From the youtube 16.11.23
func maxProfit2(prices []int) int {

	curBuy, curSell := math.MaxFloat64*-1, float64(0)

	for _, v := range prices {
		nextBuy := math.Max(curBuy, curSell-float64(v))
		nextSell := math.Max(curSell, curBuy+float64(v))
		curBuy = nextBuy
		curSell = nextSell
	}

	return int(curSell)
}

func maxProfit(prices []int) int {
	var sell int

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			sell += prices[i+1] - prices[i]
		}
	}
	return sell
}
