package main

import "fmt"

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1})) //2
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))          //3
	fmt.Println(maxDistToClosest([]int{0, 1}))                //1
	fmt.Println(maxDistToClosest([]int{0, 0, 1}))             //2
}

func maxDistToClosest(seats []int) int {
	empty := 0
	result := 0
	idxFirst, idxLast := -1, -1

	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			empty = 0
			if idxFirst == -1 {
				idxFirst = i
			}
			idxLast = i
		} else {
			empty++
			result = max(0, result, (empty+1)/2)
		}
	}
	result = max(result, idxFirst, len(seats)-1-idxLast)
	return result
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}
