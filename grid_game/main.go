package main

import "math"

func main() {
	println(gridGame([][]int{{2, 5, 4}, {1, 5, 1}})) //4
}

/*
Initialize firstRowSum with the sum of all elements in the first row of grid. Initialize secondRowSum as 0.

Set minimumSum to a very large value (LONG_LONG_MAX).

Iterate through the indices of the first row:

Subtract the current element of the first row from firstRowSum.
Calculate the maximum value between firstRowSum and secondRowSum; This would be the highest number of points the second robot can get if the first robot turns down at the current index.
Update minimumSum with the smaller value between minimumSum and the calculated maximum.
Add the current element of the second row to secondRowSum.
Return minimumSum.
*/
func gridGame(grid [][]int) int64 {
	var firstRowSum int
	secondRowSum := 0
	minimumSum := math.MaxInt

	for _, v := range grid[0] {
		firstRowSum += v
	}

	for i := 0; i < len(grid[0]); i++ {
		firstRowSum -= grid[0][i]
		maximumSum := max(firstRowSum, secondRowSum)
		minimumSum = min(minimumSum, maximumSum)
		secondRowSum += grid[1][i]
	}
	return int64(minimumSum)
}
