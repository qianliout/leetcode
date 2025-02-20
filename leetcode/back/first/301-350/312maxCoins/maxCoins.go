package main

import (
	"fmt"

	"outback/leetcode/back/common"
)

func main() {
	nums := []int{3, 1, 5, 8}
	res := maxCoins(nums)
	fmt.Println("res is ", res)
}

/*
有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
现在要求你戳破所有的气球。每当你戳破一个气球 i 时，你可以获得 nums[left] * nums[i] * nums[right] 个硬币。
这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
求所能获得硬币的最大数量。
说明:

	你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
	0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100

示例:
输入: [3,1,5,8]
输出: 167
解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []

	coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
*/
func maxCoins(nums []int) int {
	newNums := append([]int{1}, nums...)
	newNums = append(newNums, 1)
	n := len(nums)
	// dp[i][j]表法i,j之间气球的值，不包括i,j
	dp := make(map[int]map[int]int)
	// 初始化map
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
	}
	// 初值.初值都是0,所以这里可以不显式的指定
	for i := n; i >= 0; i-- {
		for j := i + 1; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = common.Max(dp[i][j], dp[i][k]+dp[k][j]+newNums[i]*newNums[k]*newNums[j])
			}
		}
	}
	return dp[0][n+1]
}
