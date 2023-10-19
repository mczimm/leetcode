package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{3, 7, 6, 4, 3, 0, 1, 0}))       //false
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))                   //true
	fmt.Println(validMountainArray([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})) //false
	fmt.Println(validMountainArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})) //false
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	left := 0
	right := len(arr) - 1

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			left = i + 1
		} else if arr[i] == arr[i+1] {
			return false
		}
	}
	for j := len(arr) - 1; j > 0; j-- {
		if arr[j] < arr[j-1] {
			right = j - 1
		} else if arr[j] == arr[j-1] {
			return false
		}
	}
	return left > 0 && right < len(arr)-1 && left == right
}
