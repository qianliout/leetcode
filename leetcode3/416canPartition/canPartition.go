package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 2, 5}))
	fmt.Println(canPartition5([]int{1, 2, 2, 5}))
	fmt.Println(canPartition5([]int{1, 2, 5}))
	fmt.Println(canPartition6([]int{1, 2, 5}))
}

func canPartition1(nums []int) bool {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if sm&1 == 1 {
		return false
	}
	// dp[i] 表示，数组和是 i 能不能等分
	dp := make(map[int]bool)
	dp[0] = true
	for _, ch := range nums {
		for j := sm >> 1; j >= 0; j-- {
			dp[j] = dp[j] || dp[j-ch]
		}
	}

	return dp[sm>>1]
}

// 01背包问题
func canPartition2(nums []int) bool {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if sm&1 == 1 {
		return false
	}
	target := sm >> 1
	//  代表考虑前i个数值(不包括 i)，其选择数字总和不超过j的最大价值。
	dp := make([]map[int]int, len(nums)+1)

	for i := range dp {
		dp[i] = make(map[int]int)
	}

	dp[0][0] = 0 // 0号之前，什么都不选，最大值当然是0了

	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]

		for j := 0; j <= target; j++ {
			no := dp[i-1][j]
			yes := 0
			if j >= ch {
				yes = dp[i-1][j-ch] + ch
			}
			dp[i][j] = Max(no, yes)
		}
	}
	return dp[len(nums)-1][target] == target
}

func canPartition(nums []int) bool {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if sm&1 == 1 {
		return false
	}
	target := sm >> 1
	//  代表考虑前i个数值，其选择数字总和不超过j的最大价值。
	dp := make([]map[int]bool, len(nums)+1)

	for i := range dp {
		dp[i] = make(map[int]bool)
	}
	// 先处理第一件物品
	// for j := 0; j <= target; j++ {
	// 	if j >= nums[0] {
	// 		dp[0][j] = true
	// 	}
	// }
	dp[0][0] = true // index 0 之前什么都不选，最大价值就是0，所以结果是 true
	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]
		for j := 0; j <= target; j++ {
			no := dp[i-1][j]
			yes := false
			if j >= ch {
				yes = dp[i-1][j-ch]
			}
			dp[i][j] = yes || no
		}
	}

	return dp[len(nums)][target]
}

// 滚动数组
func canPartition3(nums []int) bool {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if sm&1 == 1 {
		return false
	}
	target := sm >> 1
	//  代表考虑前i个数值，其选择数字总和不超过j的最大价值。
	dp := make([][]int, 2)

	for i := range dp {
		dp[i] = make([]int, target)
	}
	// 初值
	dp[0&1][0] = 1
	for i := 1; i <= len(nums); i++ {
		ch := nums[i]
		for j := 0; j <= target; j++ {
			no := dp[(i-1)&1][j]
			yes := 0
			if j >= ch {
				yes = dp[(i-1)&1][j-ch] + ch
			}
			dp[i&1][j] = Max(dp[i&1][j], no, yes)
		}
	}
	return dp[len(nums)&1][target] == target
}

// 直接使用1维数组
func canPartition4(nums []int) bool {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if sm&1 == 1 {
		return false
	}
	target := sm >> 1
	//  代表考虑前i个数值，其选择数字总和不超过j的最大价值。
	dp := make([]int, target+1)

	// 初值
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]
		for j := target; j >= 0; j-- {
			no := dp[j]
			yes := 0
			if j >= ch {
				yes = dp[j-ch] + ch
			}
			dp[j] = Max(dp[j], no, yes)
		}
	}
	return dp[target] == target
}

func canPartition5(nums []int) bool {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if sm&1 == 1 {
		return false
	}
	target := sm >> 1
	//  代表考虑前i个数值，能不能刚好选出和是 j 的集合。
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	// 初值
	dp[0][0] = true

	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]
		for j := 0; j <= target; j++ {
			no := dp[i-1][j]
			yes := false
			if j >= ch {
				yes = dp[i-1][j-ch]
			}
			dp[i][j] = dp[i][j] || no || yes
		}
	}

	return dp[len(nums)][target]
}

func canPartition6(nums []int) bool {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if sm&1 == 1 {
		return false
	}
	target := sm >> 1
	//  代表考虑前i个数值，能不能刚好选出和是 j 的集合。
	dp := make([][]bool, 2)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	// 初值
	dp[0][0] = true

	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]
		for j := 0; j <= target; j++ {
			no := dp[(i-1)&1][j]
			yes := j >= ch && dp[(i-1)&1][j-ch]
			dp[i&1][j] = dp[i&1][j] || no || yes
		}
	}
	return dp[len(nums)&1][target]
}

func canPartition7(nums []int) bool {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if sm&1 == 1 {
		return false
	}
	target := sm >> 1
	//  代表考虑前i个数值，能不能刚好选出和是 j 的集合。
	dp := make([]bool, target+1)
	// 初值
	dp[0] = true
	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]
		for j := target; j >= 0; j-- {
			no := dp[j]
			yes := j >= ch && dp[j-ch]
			dp[j] = dp[j] || no || yes
		}
	}
	return dp[target]
}
