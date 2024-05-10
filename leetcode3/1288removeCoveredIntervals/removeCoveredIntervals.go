package main

import (
	"fmt"
	"sort"
)

func main() {
	inter := [][]int{{1, 4}, {3, 6}, {2, 8}}
	fmt.Println(removeCoveredIntervals(inter))
}

func removeCoveredIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sub := 0
	sort.Sort(Intervals(intervals))
	fmt.Println(intervals)
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= start && intervals[i][1] <= end {
			sub++
		} else {
			start, end = intervals[i][0], intervals[i][1]
		}

	}
	return len(intervals) - sub
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
