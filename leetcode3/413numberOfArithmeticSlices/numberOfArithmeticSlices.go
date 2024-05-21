package main

func main() {

}

func numberOfArithmeticSlices(nums []int) int {
	// dp[i] 表示 以 nums[i]结尾的连续子数组
	dp := make([]int, len(nums))
	ans := 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
			ans += dp[i]
		} else {
			dp[i] = 0
		}
	}
	return ans
}
