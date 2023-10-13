package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14)) //6
	fmt.Println(numberOfSteps(8))  //4
}

func numberOfSteps(num int) int {
	var cnt int

	if num == 0 {
		return num
	}

	for {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
		cnt++
		if num == 0 {
			break
		}
	}

	return cnt
}
