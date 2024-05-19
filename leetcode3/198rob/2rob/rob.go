package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{1, 1}))
}

func rob2(nums []int) int {
	mem := make(map[int]int)
	// return dfs(nums, len(nums)-1, mem)
	return dfs2(nums, 0, mem)
}

// 从后向前走
func dfs(nums []int, start int, mem map[int]int) int {
	if start < 0 {
		return 0
	}
	if va, ok := mem[start]; ok {
		return va
	}
	a := dfs(nums, start-1, mem)
	b := dfs(nums, start-2, mem) + nums[start]
	if a > b {
		mem[start] = a
		return a
	}
	mem[start] = b
	return b
}

// 从前向后走
func dfs2(nums []int, start int, mem map[int]int) int {
	if start >= len(nums) {
		return 0
	}
	if va, ok := mem[start]; ok {
		return va
	}
	a := dfs2(nums, start+1, mem)
	b := dfs2(nums, start+2, mem) + nums[start]
	if a > b {
		mem[start] = a
		return a
	}
	mem[start] = b
	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp1 := make([]int, len(nums)+1) // rob
	dp2 := make([]int, len(nums)+1) // not rob

	dp1[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp1[i] = Max(dp2[i-1] + nums[i])
		dp2[i] = Max(dp1[i-1], dp2[i-1])
	}

	return Max(dp2[len(nums)-1], dp1[len(nums)-1])
}
