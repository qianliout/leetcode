package main

func main() {

}

func findTargetSumWays(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for i, num := range nums {
			dp[j] = dp[j] + dp[i-num]
		}
	}
	return dp[target]
}
