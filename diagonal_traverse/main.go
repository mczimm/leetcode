package main

import "fmt"

/*
The findDiagonalOrder function iterates over all possible diagonals (d ranges from 0 to m+n-2).

Depending on the direction (up or down), it traverses the diagonal and appends elements to the result array.

The min helper function ensures bounds are respected when calculating starting points for diagonals.

The direction alternates (up = !up) after processing each diagonal.

This solution is efficient with a time complexity of O(m * n) since each element is visited exactly once.
*/

func findDiagonalOrder(mat [][]int) []int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return []int{}
	}

	m, n := len(mat), len(mat[0])
	result := make([]int, 0, m*n)
	up := true

	for d := 0; d < m+n-1; d++ {
		if up {
			i := min(d, m-1)
			j := d - i
			for i >= 0 && j < n {
				result = append(result, mat[i][j])
				i--
				j++
			}
		} else {
			j := min(d, n-1)
			i := d - j
			for j >= 0 && i < m {
				result = append(result, mat[i][j])
				i++
				j--
			}
		}
		up = !up
	}

	return result
}

func main() {
	mat1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(findDiagonalOrder(mat1)) // Output: [1,2,4,7,5,3,6,8,9]

	mat2 := [][]int{{1, 2}, {3, 4}}
	fmt.Println(findDiagonalOrder(mat2)) // Output: [1,2,3,4]
}
