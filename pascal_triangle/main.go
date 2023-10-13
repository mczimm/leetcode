/*
Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Example 2:

Input: numRows = 1
Output: [[1]]

Constraints:

1 <= numRows <= 30
*/

package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	result := [][]int{{1}}

	for i := 1; i < numRows; i++ {
		tmp := []int{1}
		for j := 1; j < i; j++ {
			t := result[i-1][j-1] + result[i-1][j]
			tmp = append(tmp, t)
		}
		tmp = append(tmp, 1)
		result = append(result, tmp)
	}

	return result
}
