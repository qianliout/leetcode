package main

import (
	"fmt"
	"sort"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	intervals := [][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}
	fmt.Println(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return 0
	}
	sort.Sort(Intervals(intervals))
	ans := 1
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= right {
			ans++
			right = intervals[i][1]
		} else {
			// 这里取最小值是这个题目的关键
			// 要消除一个，当然要消除 right 更大的，这样后面不重复的才最多
			right = Min(right, intervals[i][1])
		}
	}
	return len(intervals) - ans
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
	}
	return vi[i][1] <= vi[j][1]
}

func (vi Intervals) Swap(i, j int) {
	vi[i], vi[j] = vi[j], vi[i]
}
