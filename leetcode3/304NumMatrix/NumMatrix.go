package main

import (
	"fmt"
)

func main() {
	// matrix := [][]int{{1, 2}, {3, 4}}
	matrix := [][]int{{-1}}
	// matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}

	nm := Constructor(matrix)
	// fmt.Println(nm.SumRegion(1, 2, 2, 4))
	fmt.Println(nm.SumRegion(0, 0, 0, 0))

}

type NumMatrix struct {
	Data   [][]int
	Prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	mu := NumMatrix{
		Data:   matrix,
		Prefix: make([][]int, len(matrix)),
	}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return mu
	}

	prefix := make([][]int, len(matrix)+1)
	for i := range prefix {
		prefix[i] = make([]int, len(matrix[0])+1)
	}
	prefix[0][0] = 0
	for i := 1; i < len(prefix); i++ {
		prefix[i][0] = prefix[i-1][0] + matrix[i-1][0]
	}
	for i := 1; i < len(prefix[0]); i++ {
		prefix[0][i] = prefix[0][i-1] + matrix[0][i-1]
	}

	for i := 1; i < len(prefix); i++ {
		for j := 1; j < len(prefix[i]); j++ {
			prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	mu.Prefix = prefix
	return mu
}

// 题目中 col 和 row 是反的
func (m *NumMatrix) SumRegion(col1 int, row1 int, col2 int, row2 int) int {
	ans := m.Prefix[col2+1][row2+1] + m.Prefix[col1][row1] - m.Prefix[col1][row2+1] - m.Prefix[col2+1][row1]
	return ans
}
