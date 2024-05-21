package main

func main() {

}

// 子数组 是数组中的一个连续序列。
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	dp := make([]int, len(nums))
	ans := 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
		ans += dp[i]
	}
	return ans
}
