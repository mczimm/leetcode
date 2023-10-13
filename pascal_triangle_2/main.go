package main

import "fmt"

func main() {
	fmt.Println(getRow(3)) //{1,3,3,1}
}

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)

	if rowIndex == 0 {
		return []int{1}
	}

	for i := 0; i <= rowIndex; i++ {
		res[i] = 1
		for j := i - 1; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}
