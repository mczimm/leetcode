package main

import "fmt"

func main() {
	fmt.Println(maxFrequencyElements([]int{1, 2, 2, 3, 1, 4}))         //4
	fmt.Println(maxFrequencyElements([]int{1, 2, 3, 4, 5}))            //5
	fmt.Println(maxFrequencyElements([]int{10, 12, 11, 9, 6, 19, 11})) //2
}

func maxFrequencyElements(nums []int) int {
	//https://leetcode.com/problems/count-elements-with-maximum-frequency/editorial/
	var resFreq, currFreq, maxFreq int
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
		currFreq = freq[nums[i]]

		if currFreq > maxFreq {
			resFreq = currFreq
			maxFreq = currFreq
		} else if currFreq == maxFreq {
			resFreq += currFreq
		}
	}
	return resFreq
}
