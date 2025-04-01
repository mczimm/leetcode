package main

import "fmt"

func main() {
	//fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))      //[1,2]
	//fmt.Println(topKFrequent([]int{1, 2}, 2))                  //[1,2]
	//fmt.Println(topKFrequent([]int{-1, -1}, 1))                //[-1]
	//fmt.Println(topKFrequent([]int{3, 0, 1, 0}, 1))            //[0]
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2)) //[2,-1]
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == k {
		return nums
	}

	d := make(map[int]int)
	var res []int
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		d[num]++
	}
	for num, count := range d {
		freq[count] = append(freq[count], num)
	}

	for i := len(freq) - 1; i >= 0; i-- {
		if len(res) < k && freq[i] != nil {
			res = append(res, freq[i]...)
		}
	}

	return res
}
