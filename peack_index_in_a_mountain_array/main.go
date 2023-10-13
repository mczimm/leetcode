// An array arr a mountain if the following properties hold:

// arr.length >= 3
// There exists some i with 0 < i < arr.length - 1 such that:
// arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
// arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
// Given a mountain array arr, return the index i such that arr[0] < arr[1] < ... < arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].

// You must solve it in O(log(arr.length)) time complexity.

package main

import "fmt"

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 10, 11, 5, 2}))
	fmt.Println(peakIndexInMountainArray([]int{7, 2, 3, 1, 2}))
}

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := (left + right) / 2

		if arr[middle] < arr[middle+1] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}
