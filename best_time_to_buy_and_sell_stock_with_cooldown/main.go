package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2})) // 3
}

// From the youtube 16.11.23
func maxProfit(prices []int) int {
	curBuy, curSell, coolDown := math.MaxFloat64*-1, float64(0), float64(0)

	for _, v := range prices {
		buy := math.Max(curBuy, coolDown-float64(v))
		sell := curBuy + float64(v)
		coolDown = math.Max(curSell, coolDown)

		curBuy = buy
		curSell = sell
	}
	return int(math.Max(curSell, coolDown))
}
