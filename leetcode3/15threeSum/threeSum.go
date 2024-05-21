package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > 0 && nums[j] == nums[j-1] {
				continue
			}
			dfs(nums, i, j, j+1, &ans)
		}
	}
	return ans
}

func dfs(nums []int, a, b, start int, ans *[][]int) {
	if start >= len(nums) {
		return
	}
	if nums[a]+nums[b]+nums[start] == 0 {
		*ans = append(*ans, []int{nums[a], nums[b], nums[start]})
	}
	dfs(nums, a, b, start+1, ans)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			tag := nums[l] + nums[r] + nums[i]
			if tag > 0 {
				r--
			} else if tag < 0 {
				l++
			} else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	return ans
}
