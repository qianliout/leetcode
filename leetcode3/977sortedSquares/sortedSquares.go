package main

import (
	"sort"
)

func main() {

}

func sortedSquares(nums []int) []int {
	ans := make([]int, 0)
	for i := range nums {
		ans = append(ans, nums[i]*nums[i])
	}

	sort.Ints(ans)
	return ans
}
