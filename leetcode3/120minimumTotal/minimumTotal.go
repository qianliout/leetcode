package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
	fmt.Println(minimumTotal([][]int{{-1}, {2, 3}, {1, -1, -3}}))
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 && len(triangle[0]) == 1 {
		return triangle[0][0]
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = Min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}

	return triangle[0][0]
}
