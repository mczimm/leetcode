package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)) //true
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)) //false
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, 0)

	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n -= 1
		}
	}
	return n <= 0
}
