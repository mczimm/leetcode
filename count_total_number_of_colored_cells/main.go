package main

import "fmt"

func main() {
	fmt.Println(coloredCells(5))
}

func coloredCells(n int) int64 {
	//using the formula 𝐶(𝑡)=2𝑡2−2𝑡+1
	return int64(2*n*n - 2*n + 1)
}
