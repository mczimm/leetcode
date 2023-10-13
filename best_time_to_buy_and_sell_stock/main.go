package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))    //5
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))       //0
	fmt.Println(maxProfit([]int{1, 2}))                //1
	fmt.Println(maxProfit([]int{2, 4, 1}))             //2
	fmt.Println(maxProfit([]int{2, 1, 2, 0, 1}))       //1
	fmt.Println(maxProfit([]int{2, 1, 2, 1, 0, 1, 2})) //2
	fmt.Println(maxProfit([]int{3, 2, 6, 5, 0, 3}))    //4
}

func maxProfit(prices []int) int {
	sell := 0
	buy := 10000
	pos := 0
	res := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < buy && i+1 < len(prices) {
			buy = prices[i]
			sell = 0
			pos = i
			continue
		}
		if prices[i] > sell {
			if pos < i {
				sell = prices[i]
			}
		}
		if res < sell-buy {
			res = sell - buy
		}
	}

	return res
}
