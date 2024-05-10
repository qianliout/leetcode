package main

import (
	"fmt"
)

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
	// [8,1,5,4,3,2,1,1,0,0]
}

func dailyTemperatures(temperatures []int) []int {
	stark := make([]int, 0)
	ans := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		// 递减栈
		// 注意，这里求下一个最近的，所以是<=
		for len(stark) > 0 && temperatures[stark[len(stark)-1]] <= temperatures[i] {
			stark = stark[:len(stark)-1]
		}
		if len(stark) == 0 {
			ans[i] = 0
		} else {
			ans[i] = stark[len(stark)-1] - i
		}
		stark = append(stark, i)
	}
	return ans
}
