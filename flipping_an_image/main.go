package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}})) //[[1,0,0],[0,1,0],[1,1,1]]
}

func flipAndInvertImage(image [][]int) [][]int {
	var res [][]int

	for i := 0; i < len(image); i++ {
		var tmp []int
		for j := len(image[i]) - 1; j >= 0; j-- {
			if image[i][j] == 0 {
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, 0)
			}
		}
		res = append(res, tmp)
	}
	return res
}
