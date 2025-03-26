package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) //[1,2,3,6,9,8,7,4,5]
}

func spiralOrder(matrix [][]int) []int {
	dm := make(map[int]int)
	dm[0] = 1
	dm[1] = 0
	dm[1] = 0
	dm[0] = -1
	dm[0] = -1
	dm[-1] = 0
	dm[-1] = 0
	dm[0] = 1

	lstr := len(matrix)
	lcol := len(matrix[0])
	r, c := 0, 0
	direction := make(map[int]int)
	direction[0] = 1
	var ans []int
	var dk int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			ans = append(ans, matrix[i][j])
			matrix[i][j] = 101
			rs, cs := dk, direction[dk]
			nr, nc := r+rs, c+cs
			if (nr >= 0 && lstr > nr) && (nc >= 0 && lcol > nc) && matrix[i][j] <= 100 {
				r, c = nr, nc
			} else {

			}
		}
	}

}
