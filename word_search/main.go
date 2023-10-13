package main

import "fmt"

func main() {
	one := []byte("ABCE")
	two := []byte("SFCS")
	three := []byte("ADEE")
	board := make([][]byte, 0)
	board = append(board, one, two, three)
	word := "ABCCED"
	//word := "BFD"

	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && backtrack(board, visited, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func backtrack(board [][]byte, visited [][]bool, word string, i, j, idx int) bool {
	if idx == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || visited[i][j] || board[i][j] != word[idx] {
		return false
	}

	visited[i][j] = true

	if backtrack(board, visited, word, i+1, j, idx+1) ||
		backtrack(board, visited, word, i-1, j, idx+1) ||
		backtrack(board, visited, word, i, j+1, idx+1) ||
		backtrack(board, visited, word, i, j-1, idx+1) {
		return true
	}

	visited[i][j] = false
	return false
}
