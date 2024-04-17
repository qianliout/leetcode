package main

import (
	"fmt"
	"sort"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	fmt.Println(maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}}))
	fmt.Println(maxEnvelopes([][]int{{4, 5}, {6, 7}, {2, 3}}))
}

// 可以得到结果，但是超时了
func maxEnvelopes(envelopes [][]int) int {
	sort.Sort(Enve(envelopes))
	// dp[i] 表示 envelopes[i]时，可以装多个信封
	dp := make(map[int]int)
	ans := 0
	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if contain(envelopes[i], envelopes[j]) {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		ans = Max(ans, dp[i])
	}
	return ans
}

type Enve [][]int

func (vi Enve) Len() int {
	return len(vi)
}

func (vi Enve) Less(i, j int) bool {
	if vi[i][0] < vi[j][0] {
		return true
	} else if vi[i][0] > vi[j][0] {
		return false
	} else {
		return vi[i][1] < vi[j][1]
	}
}

func (vi Enve) Swap(i, j int) {
	vi[i], vi[j] = vi[j], vi[i]
}

func contain(b, a []int) bool {
	return a[0] < b[0] && a[1] < b[1]
}
