package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i, ch := range nums {
		if i > 0 && nums[i-1] == ch {
			continue
		}

		ts := threeSum(nums[i+1:], target-ch)
		for j := range ts {
			ans = append(ans, append(ts[j], ch))
		}
	}
	return ans
}

func twoSum2(nums []int, target int) []int {
	exit := make(map[int]int)
	for i, ch := range nums {
		if in, ok := exit[target-ch]; ok {
			return []int{i, in}
		}
		exit[ch] = i
	}
	ans := make([]int, 0)
	return ans
}

func twoSum(nums []int, target int) []int {
	exit := make(map[int]int)
	for i, ch := range nums {
		exit[ch] = i
	}
	for i, ch := range nums {
		if in, ok := exit[target-ch]; ok && in != i {
			return []int{i, in}
		}
	}
	return []int{}
}

func threeSum(nums []int, target int) [][]int {
	// nums 中 有重复数据,所以就得排序
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i, ch := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			su := ch + nums[left] + nums[right]
			if su > target {
				right--
			} else if su < target {
				left++
			} else if su == target {
				ans = append(ans, []int{ch, nums[left], nums[right]})

				// for left+1 < right && nums[left] == nums[left+1] {
				// 上面这种写法也是对的，但是推荐这种写法
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right-1 && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return ans
}
