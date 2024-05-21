package main

func main() {

}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i]
		if dp[i-1]+nums[i] > dp[i] {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
