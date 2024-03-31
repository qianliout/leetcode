package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	left, right, up, down, used := 0, n-1, 0, n-1, 0
	for {
		for i := left; i <= right; i++ {
			used++
			ans[up][i] = used
		}
		up++
		for i := up; i <= down; i++ {
			used++
			ans[i][right] = used
		}
		right--
		for i := right; i >= left; i-- {
			used++
			ans[down][i] = used
		}
		down--
		for i := down; i >= up; i-- {
			used++
			ans[i][left] = used
		}
		left++

		if used >= n*n {
			break
		}
	}
	return ans
}
