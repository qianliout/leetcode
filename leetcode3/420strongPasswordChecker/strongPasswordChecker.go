package main

func main() {

}

func strongPasswordChecker(password string) int {
	n := len(password)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 0
	return 0

}
