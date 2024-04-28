package main

import (
	"sort"
)

func main() {
}

func findMinArrowShots(points [][]int) int {
	return sortByEnd(points)
}

func sortByEnd(points [][]int) int {
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
		}
	}
	return ans
}

type Points [][]int

func (vi Points) Len() int {
	return len(vi)
}

func (vi Points) Less(i, j int) bool {
	if vi[i][1] < vi[j][1] {
		return true
	}
	if vi[i][1] > vi[j][1] {
		return false
	}
	return vi[i][0] <= vi[j][0]
}

func (vi Points) Swap(i, j int) {
	vi[i], vi[j] = vi[j], vi[i]
}
