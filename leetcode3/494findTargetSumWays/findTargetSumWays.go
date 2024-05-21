package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays3([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays4([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays4([]int{100}, -200))
	fmt.Println(findTargetSumWays5([]int{200}, -200))
	fmt.Println(findTargetSumWays6([]int{100}, -200))
	fmt.Println(findTargetSumWays5([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays2(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for i, num := range nums {
			dp[j] = dp[j] + dp[i-num]
		}
	}
	return dp[target]
}

// https://mp.weixin.qq.com/s?__biz=MzU4NDE3MTEyMA==&mid=2247488724&idx=1&sn=68b106ec37730b9ce3988195ae45ac7b
func findTargetSumWays3(nums []int, target int) int {
	dp := make([]map[int]int, len(nums)+1)
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	sm := 0
	for _, ch := range nums {
		sm += ch
	}

	sub := (sm - target) / 2
	if sm < target || (sm-target)&1 == 1 {
		return 0
	}

	dp[0][0] = 1

	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]
		for j := 0; j <= sub; j++ {
			no := dp[i-1][j]
			yes := dp[i-1][j-ch]
			if j-ch < 0 {
				yes = 0
			}

			dp[i][j] += no + yes
		}
	}
	return dp[len(nums)][sub]
}

func findTargetSumWays(nums []int, target int) int {
	mem := make(map[string]int)
	return dfs(nums, target, 0, 0, mem)
}

func dfs(nums []int, target, start, cur int, mem map[string]int) int {
	key := fmt.Sprintf("%d_%d", start, cur)
	if val, ok := mem[key]; ok {
		return val
	}

	if start >= len(nums) {
		if target == cur {
			mem[key] = 1
		}
		return mem[key]
	}
	add := dfs(nums, target, start+1, cur+nums[start], mem)
	sub := dfs(nums, target, start+1, cur-nums[start], mem)
	mem[key] = add + sub

	return add + sub
}

// 注意target 可能是负数
func findTargetSumWays4(nums []int, target int) int {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if (target+sm)&1 != 0 {
		return 0
	}

	find := (target + sm) >> 1

	// 这里如果写成常规的数组，则需要考虑负数导致的下标问题
	dp := make([]map[int]int, len(nums)+1)
	for i := range dp {
		dp[i] = make(map[int]int, find+1)
	}

	dp[0][0] = 1

	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]
		for j := find; j >= 0; j-- {
			no := dp[i-1][j]
			yes := 0
			if j >= ch {
				yes = dp[i-1][j-ch]
			}
			dp[i][j] = yes + no
		}
	}
	return dp[len(nums)][find]
}

// 注意target 可能是负数
func findTargetSumWays5(nums []int, target int) int {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}

	if (target+sm)&1 != 0 {
		return 0
	}

	find := (target + sm) >> 1

	dp := make(map[int]int, find+1)
	dp[0] = 1

	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]
		for j := find; j >= 0; j-- {
			no := dp[j]
			yes := 0
			if j >= ch {
				yes = dp[j-ch]
			}
			dp[j] = no + yes
		}
	}
	return dp[find]
}

func findTargetSumWays6(nums []int, target int) int {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if (target+sm)&1 != 0 {
		return 0
	}

	find := (target + sm) >> 1
	// find 表示找到的组合相加，可能组成 find 的值，因为target 有负数，这里的 find 可能是一个负数,最好的方式就是直接使用 map，
	dp := make([][]int, len(nums)+1)

	for i := range dp {
		dp[i] = make([]int, find+1)
	}

	dp[0][0] = 1

	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]
		for j := 0; j <= find; j++ {
			no := dp[i-1][j]
			yes := 0
			if j >= ch {
				yes = dp[i-1][j-ch]
			}
			dp[i][j] = yes + no

		}
	}
	return dp[len(nums)][find]
}
