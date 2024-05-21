package main

import (
	"sort"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
}

func findMinArrowShots(points [][]int) int {
	return sortByStart(points)
}

// 以起始点排序
func sortByStart(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Sort(Points(points))
	end := points[0][1]
	ans := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			end = points[i][1]
			ans++
		} else {
			end = Min(end, points[i][1])
		}
	}
	return ans
}

type Points [][]int

func (vi Points) Len() int {
	return len(vi)
}

func (vi Points) Less(i, j int) bool {
	if vi[i][0] < vi[j][0] {
		return true
	}
	if vi[i][0] > vi[j][0] {
		return false
	}
	return vi[i][1] <= vi[j][1]
}

func (vi Points) Swap(i, j int) {
	vi[i], vi[j] = vi[j], vi[i]
}
