package main

func main() {
	println(checkValid([][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}})) //true
	println(checkValid([][]int{{1, 1, 1}, {1, 2, 3}, {1, 2, 3}})) //false
}

func checkValid(matrix [][]int) bool {
	//rows := make([]int, len(matrix))

	for _, row := range matrix {
		cols := [101]int{}
		for _, col := range row {
			if cols[col] != 0 {
				return false
			}
			cols[col] = col
		}
	}
	for i := range len(matrix) {
		rows := [101]int{}
		for j := range len(matrix[0]) {
			val := matrix[j][i]
			if rows[val] != 0 {
				return false
			}
			rows[val] = val
		}
	}
	//for i := 0; i < len(matrix); i++ {
	//	for j := 0; j < len(matrix[0]); j++ {
	//		rows[matrix[i][j]]++
	//	}
	//}
	return true
}
