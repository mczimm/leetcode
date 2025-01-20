package main

func main() {
	println(firstCompleteIndex([]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}}))                                 //2
	println(firstCompleteIndex([]int{2, 8, 7, 4, 1, 3, 5, 6, 9}, [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}})) //3
	println(firstCompleteIndex([]int{1, 4, 5, 2, 6, 3}, [][]int{{4, 3, 5}, {1, 2, 6}}))                     //1
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	nRows, nCols := len(mat), len(mat[0])
	rowCounts := make([]int, nRows)
	colCounts := make([]int, nCols)
	positions := make(map[int][2]int)
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			positions[mat[i][j]] = [2]int{i, j}
		}
	}
	for i, v := range arr {
		row, col := positions[v][0], positions[v][1]
		rowCounts[row]++
		colCounts[col]++

		if rowCounts[row] == nCols || colCounts[col] == nRows {
			return i
		}
	}
	return -1
}

func find(num int, mat [][]int) (int, int) {
	for i, row := range mat {
		for j, val := range row {
			if num == val {
				return i, j
			}
		}
	}
	return -1, -1
}
