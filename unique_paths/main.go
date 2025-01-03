package main

func main() {
	println(uniquePaths(3, 7)) //28
	println(uniquePaths(3, 2)) //3
}

//recursion
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

//dp
func uniquePaths2(m int, n int) int {
	dp := [100][100]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
