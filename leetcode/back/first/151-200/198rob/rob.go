package main

import (
	"fmt"

	"qianliout/leetcode/leetcode/common/utils"
)

func main() {
	nums := []int{1, 2, 3, 1}
	res := rob(nums)
	fmt.Println("res is ", res)
}

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
示例 1:
输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2:
输入: [2,7,9,3,1]
输出: 12
解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp0 := make(map[int]int) // 表示不偷
	dp1 := make(map[int]int) // 表示偷

	dp0[0] = 0
	dp1[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp0[i] = utils.Max(dp0[i-1], dp1[i-1])
		dp1[i] = dp0[i-1] + nums[i]
	}
	return utils.Max(dp0[len(nums)-1], dp1[len(nums)-1])
}
