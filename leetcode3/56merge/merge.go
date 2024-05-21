package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return intervals
	}
	sort.Sort(Inter(intervals))
	start, end := intervals[0][0], intervals[0][1]
	ans := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			ans = append(ans, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else {
			if end < intervals[i][1] {
				end = intervals[i][1]
			}
		}
	}
	ans = append(ans, []int{start, end})
	return ans
}

type Inter [][]int

func (vi Inter) Len() int {
	return len(vi)
}

func (vi Inter) Less(i, j int) bool {
	if vi[i][0] < vi[j][0] {
		return true
	} else if vi[i][0] > vi[j][0] {
		return false
	} else {
		return vi[i][1] < vi[j][1]
	}
}

func (vi Inter) Swap(i, j int) {
	vi[i], vi[j] = vi[j], vi[i]
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)
}
