package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	col := 0
	row := n - 1
	for col < m && row >= 0 {
		ma := matrix[col][row]
		if ma == target {
			return true
		}
		if ma > target {
			row--
			continue
		}
		if ma < target {
			col++
		}
	}

	return false
}
