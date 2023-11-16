package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2)) //8
}

// From youtube 16.11.23
func maxProfit(prices []int, fee int) int {
	curBuy, curSell := math.MaxFloat64*-1, float64(0)

	for _, v := range prices {
		nextBuy := math.Max(curBuy, curSell-float64(v+fee))
		nextSell := math.Max(curSell, curBuy+float64(v))
		curBuy = nextBuy
		curSell = nextSell
	}

	return int(curSell)
}
