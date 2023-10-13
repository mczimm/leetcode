package main

import "fmt"

func main() {
	fmt.Println(countBits(2)) //[0,1,1]
}

func countBits(n int) []int {
	ans := []int{0}

	for i := 1; i < n+1; i++ {
		ans = append(ans, ans[i>>1]+i&1)
	}
	return ans
}
