/*
A peak element in a 2D grid is an element that is strictly greater than all of its adjacent neighbors to the left, right, top, and bottom.

Given a 0-indexed m x n matrix mat where no two adjacent cells are equal, find any peak element mat[i][j] and return the length 2 array [i,j].

You may assume that the entire matrix is surrounded by an outer perimeter with the value -1 in each cell.

You must write an algorithm that runs in O(m log(n)) or O(n log(m)) time.

Input: mat = [[1,4],[3,2]]
Output: [0,1]
Explanation: Both 3 and 4 are peak elements so [1,0] and [0,1] are both acceptable answers.

Input: mat = [[10,20,15],[21,30,14],[7,16,32]]
Output: [1,1]
Explanation: Both 30 and 32 are peak elements so [1,1] and [2,2] are both acceptable answers.

No two adjacent cells are equal.
*/

package main

import "fmt"

func main() {
	fmt.Println(findPeakGrid([][]int{{1, 4}, {3, 2}}))                                                 // [0,1]
	fmt.Println(findPeakGrid([][]int{{11, 27, 32, 31, 14}, {26, 4, 11, 25, 1}, {25, 17, 30, 19, 28}})) // [0,2]
	fmt.Println(findPeakGrid([][]int{{7, 2, 3, 1, 2}, {6, 5, 4, 2, 1}}))                               // [0,0]
	fmt.Println(findPeakGrid([][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}))                        // [2,2]
	fmt.Println(findPeakGrid([][]int{{10, 50, 40, 30, 20}, {1, 500, 2, 3, 4}}))                        // [1,1]
	fmt.Println(findPeakGrid([][]int{{10, 30, 40, 50, 20}, {1, 3, 2, 500, 4}}))                        // [1,3]
}

func findPeakGrid(mat [][]int) []int {
	n, m, max, peak := 0, 0, 0, 0

	for i := 0; i < len(mat); i++ {
		peak = findPeakElementBrootForce(mat[i])
		if max < mat[i][peak] {
			max = mat[i][peak]
			n = i
			m = peak
		}
	}
	return []int{n, m}
}

func findPeakElementBrootForce(nums []int) int {
	res := -1
	ind := 0
	for i, v := range nums {
		if v > res {
			res = v
			ind = i
		}
	}
	return ind
}
