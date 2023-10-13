/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.
*/

package main

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}} //[[1,0,1],[0,0,0],[1,0,1]]
	//matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}} //[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
	setZeroes(matrix)
}

func setZeroes(matrix [][]int) {
	var idx []int
	line := make(map[int]struct{})
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				idx = append(idx, j)
				line[i] = struct{}{}
			}
		}
	}

	//cols
	for i := 0; i < len(matrix); i++ {
		for _, v := range idx {
			matrix[i][v] = 0
		}
	}

	//rows
	for i := range line {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = 0
		}
	}
}
