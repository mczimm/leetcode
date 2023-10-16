package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))  //[1,2]
	fmt.Println(twoSum([]int{-1, 0}, -1))        //[1,2]
	fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9)) //[1,2]
	fmt.Println(twoSum2([]int{-1, 0}, -1))       //[1,2]
}

// My first attempt
func twoSum(numbers []int, target int) []int {
	var res []int
	numbers = append([]int{0}, numbers...)

	for i := 1; i < len(numbers); i++ {
		for j := 1; j < len(numbers); j++ {
			if i == j {
				continue
			}
			if numbers[i]+numbers[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}

// My second attempt
func twoSum2(numbers []int, target int) []int {
	var res []int
	numbers = append([]int{0}, numbers...)
	left, right := 1, len(numbers)-1

	for i := 1; i < len(numbers); i++ {
		if left == right {
			continue
		}
		if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			res = append(res, left, right)
			break
		}
	}
	return res
}
