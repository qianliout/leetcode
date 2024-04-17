package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3}))
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8}))
	fmt.Println(largestDivisibleSubset([]int{3, 4, 16, 8}))
	fmt.Println(largestDivisibleSubset([]int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}))
}

func largestDivisibleSubset(nums []int) []int {
	// dp[i] 表示index是 i 的数最大整除子集
	ans := make([]int, 0)
	if len(nums) <= 1 {
		return nums
	}
	sort.Ints(nums)
	dp := make([][]int, len(nums)+1)
	for i := range nums {
		dp[i] = make([]int, 0)
	}

	dp[0] = []int{}
	for i := 0; i < len(nums); i++ {
		lev := make([]int, 0)
		for j := i - 1; j >= 0; j-- {
			if (nums[i]%nums[j] == 0 || nums[j]%nums[i] == 0) && len(dp[j]) > len(lev) {
				lev = dp[j]
			}
		}
		dp[i] = append(append([]int{}, lev...), nums[i])
		if len(dp[i]) > len(ans) {
			ans = dp[i]
		}
	}
	return ans
}
