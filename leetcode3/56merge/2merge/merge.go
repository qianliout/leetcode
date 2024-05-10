package main

import (
	"sort"

	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return intervals
	}
	sort.Sort(Intervals(intervals))
	start, end := intervals[0][0], intervals[0][1]
	ans := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			ans = append(ans, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else {
			end = Max(end, intervals[i][1])
		}
	}
	ans = append(ans, []int{start, end})
	return ans
}

type Intervals [][]int

func (vi Intervals) Len() int {
	return len(vi)
}

func (vi Intervals) Less(i, j int) bool {
	if vi[i][0] < vi[j][0] {
		return true
	} else if vi[i][0] > vi[j][0] {
		return false
	} else {
		return vi[i][1] <= vi[j][1]
	}
}

func (vi Intervals) Swap(i, j int) {
	vi[i], vi[j] = vi[j], vi[i]
}
