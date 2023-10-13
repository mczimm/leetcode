package main

import "fmt"

func main() {
	// mat := [][]int{
	// 	{1, 2, 3},1,2,3
	// 	{4, 5, 6},1,2,3
	// 	{7, 8, 9}}1,2,3
	mat := fillSame()
	// mat := fillSeq()

	fmt.Println(matrixSumDiag(mat))             //30, 12
	fmt.Println(matrixSumDiagWithoutCross(mat)) //25, 10
}

func matrixSumDiag(mat [][]int) int {
	res := 0
	left := 0
	right := len(mat) - 1

	for i := 0; i < len(mat); i++ {
		res += mat[i][left]
		res += mat[i][right]
		left++
		right--
	}

	return res
}

func matrixSumDiagWithoutCross(mat [][]int) int {
	res := 0
	left := 0
	right := len(mat) - 1

	for i := 0; i < len(mat); i++ {
		if mat[i][left] != mat[i][right] {
			res += mat[i][left]
			res += mat[i][right]
		} else {
			res += mat[i][left]
		}
		left++
		right--
	}

	return res
}

func fillSame() [][]int {
	var res [][]int
	sl := []int{1, 2, 3}
	for i := 0; i < 3; i++ {
		res = append(res, sl)
	}
	return res
}

func fillSeq() [][]int {
	res := make([][]int, 3)
	cnt := 1

	for i := 0; i < 3; i++ {
		res[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			res[i][j] = cnt
			cnt++
		}
	}

	return res
}
