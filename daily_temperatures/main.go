package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures2([]int{73, 74, 75, 71, 69, 72, 76, 73})) // [1, 1, 4, 2, 1, 1, 0, 0]
	fmt.Println(dailyTemperatures2([]int{30, 40, 50, 60}))                 // [1,1,1,0]
}

// time limit exceeded
func dailyTemperatures(temperatures []int) []int {
	var res []int

	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[i] < temperatures[j] {
				res = append(res, j-i)
				break
			}
			if j == len(temperatures)-1 {
				res = append(res, 0)
			}
			continue
		}
	}
	res = append(res, 0)
	return res
}

//not mine
func dailyTemperatures2(temps []int) []int {
	results := make([]int, len(temps))
	stack := make([]int, 0)
	/// UPVOTE !
	for i, temp := range temps {
		for len(stack) > 0 && temps[stack[len(stack)-1]] < temp {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			results[index] = i - index
		}
		stack = append(stack, i)
	}

	return results
}
