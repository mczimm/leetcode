package main

import "fmt"

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))                    //true
	fmt.Println(checkIfExist([]int{3, 1, 7, 11}))                    //false
	fmt.Println(checkIfExist([]int{0, 0}))                           //true
	fmt.Println(checkIfExist([]int{-2, 0, 10, -19, 4, 6, -8}))       //false
	fmt.Println(checkIfExist([]int{7, 1, 14, 11}))                   //true
	fmt.Println(checkIfExist([]int{-20, 8, -6, -14, 0, -19, 14, 4})) //true
}

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			if arr[i] == arr[j]*2 {
				fmt.Println(arr[i], arr[j])
				return true
			}
		}
	}
	return false
}
