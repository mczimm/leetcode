package main

import (
	"fmt"
)

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)) //[true,true,true,false,true]
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)) //[true,false,false,false,false]
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))    //[true,false,true]
}

//My first attempt
func kidsWithCandies(candies []int, extraCandies int) []bool {
	var res []bool
	var maxCandie int

	for i := 0; i < len(candies); i++ {
		if candies[i] > maxCandie {
			maxCandie = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxCandie {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
