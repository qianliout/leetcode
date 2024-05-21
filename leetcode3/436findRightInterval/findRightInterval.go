package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findRightInterval([][]int{{1, 2}}))
	fmt.Println(findRightInterval([][]int{{1, 1}, {3, 4}}))
}

func findRightInterval(intervals [][]int) []int {

	idxM := make(map[int]int)
	for i, ch := range intervals {
		idxM[ch[0]] = i
	}

	ans := make([]int, len(intervals))

	sort.Sort(Intervals(intervals))

	for i := 0; i < len(intervals); i++ {
		ans[idxM[intervals[i][0]]] = -1
		if intervals[i][0] == intervals[i][1] {
			ans[idxM[intervals[i][0]]] = idxM[intervals[i][0]]
			continue
		}
		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] >= intervals[i][1] {
				ans[idxM[intervals[i][0]]] = idxM[intervals[j][0]]
				break
			}
		}
	}
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
	}
	return vi[i][1] < vi[j][1]
}

func (vi Intervals) Swap(i, j int) {
	vi[i], vi[j] = vi[j], vi[i]
}
