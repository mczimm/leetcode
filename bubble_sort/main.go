package main

import "fmt"

func main1() {
	origArr := [4]int{4, 2, 1, 3}
	// fmt.Printf("The original array: %v\n", origArr)

	for i := 0; i < len(origArr)-1; i++ {
		for j := 0; j < len(origArr)-i-1; j++ {
			// fmt.Printf("Checking the %d\n", origArr[j])
			if origArr[j] > origArr[j+1] {
				// fmt.Printf("The %d is bigger than %d, moving to the right\n", origArr[j], origArr[j+1])
				origArr[j], origArr[j+1] = origArr[j+1], origArr[j]
			}
		}
		// fmt.Printf("Checks for current element has ended\n")
	}
	fmt.Printf("The sorted array: %v\n", origArr)
}

func main() {
	origArr := [4]int{4, 2, 1, 3}
	cnt := len(origArr)
	for cnt > 0 {
		for i := 0; i < len(origArr)-1; i++ {
			if origArr[i] > origArr[i+1] {
				origArr[i], origArr[i+1] = origArr[i+1], origArr[i]
			}
			cnt--
		}
	}
	fmt.Println(origArr)
}
