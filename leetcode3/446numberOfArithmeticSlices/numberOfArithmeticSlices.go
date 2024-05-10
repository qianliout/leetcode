package main

import (
	"fmt"
)

func main() {
	// fmt.Println(numberOfArithmeticSlices([]int{7, 7, 7, 7, 7}))
	fmt.Println(numberOfArithmeticSlices([]int{0, 1, 2, 3, 4}))
}

func numberOfArithmeticSlices(nums []int) int {
	// dp[i][j] 选中nums[i]这个数时，差值为 j 的数列个数
	// 注意，当只有两个数时，也会计算在内
	// 假如原数组是[0,1,2,3,4]，这里为了方便，数组组和下标值取一样，方便理解
	// dp[0] = {}
	// dp[1] = {1:1} 间距是1有一个元素
	// dp[2] = {1:2,2:1} 间距是1有2元素,间距是2有1元素
	// dp[3] = {1:3,2:1,3:1}
	// dp[4] = {[1:4 2:2 3:1 4:1}

	dp := make([]map[int]int, len(nums))
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			dp[i][diff] += dp[j][diff] + 1
			// 假如dp[j][diff]=1，说明在 j 这里，间距是 diff 的前面有1个元素，加上j和i，就正好是三个元素，所以可以算做结果
			// 假如dp[j][diff]=0，说明在 j 这里，间距是 diff 的前面有0个元素，加上j和i，就正好是2个元素，不能算做结果，但是 dp[j][diff]=0,对结果没有影响
			ans += dp[j][diff]
		}
	}

	return ans
}
