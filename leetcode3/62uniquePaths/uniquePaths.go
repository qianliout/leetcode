package main

func main() {

}

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := range dp {
		dp[i][0] = 1
	}
	for j := range dp[0] {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]

		}
	}
	return dp[m-1][n-1]
}
