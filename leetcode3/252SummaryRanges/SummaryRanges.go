package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(GetIntervals([]int{1, 2, 3}))
	fmt.Println(GetIntervals([]int{1, 3}))
	fmt.Println(GetIntervals([]int{1}))
}

type SummaryRanges struct {
	Data []int
}

func Constructor() SummaryRanges {
	return SummaryRanges{Data: make([]int, 0)}
}

func (this *SummaryRanges) AddNum(value int) {
	this.Data = append(this.Data, value)
	sort.Ints(this.Data)

}

func (this *SummaryRanges) GetIntervals() [][]int {
	return GetIntervals(this.Data)
}

func GetIntervals(data []int) [][]int {
	// 默认已排序好
	ans := make([][]int, 0)
	if len(data) == 0 {
		return ans
	}
	start, end := data[0], data[0]
	for i := 0; i < len(data); i++ {
		if data[i] > end+1 {
			ans = append(ans, []int{start, end})
			start, end = data[i], data[i]
		} else {
			end = data[i]
		}
	}
	ans = append(ans, []int{start, end})

	return ans
}
