package main

import "fmt"

func main() {
	fmt.Println(quickSort([]int{5, 3, 4, 2, 1}))
}

func quickSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}
	var left []int
	var right []int
	pivot := values[len(values)-1]

	for _, v := range values[0 : len(values)-1] {
		if v < pivot {
			left = append(left, v)
			continue
		}
		right = append(right, v)
	}

	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)

	res := append(sortedLeft, pivot)
	res = append(res, sortedRight...)

	return res
}
