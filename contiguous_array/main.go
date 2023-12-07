package main

import "fmt"

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))                //2
	fmt.Println(findMaxLength([]int{0, 0, 1, 1}))          //4
	fmt.Println(findMaxLength([]int{0, 0}))                //0
	fmt.Println(findMaxLength([]int{1, 1}))                //0
	fmt.Println(findMaxLength([]int{0, 0, 0, 1, 1, 1, 0})) //6
}

func findMaxLength(nums []int) int {
	var numsM = make(map[int]int)
	var counter, res int
	//some magic
	numsM[0] = -1

	for i, v := range nums {
		if v == 0 {
			counter--
		} else {
			counter++
		}
		if _, ok := numsM[counter]; !ok {
			numsM[counter] = i
		}
		//maximum between current res and current index minus saved index for current counter
		res = max(res, i-numsM[counter])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
