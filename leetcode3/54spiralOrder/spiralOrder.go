package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ans
	}

	left, right, up, down, all := 0, len(matrix[0])-1, 0, len(matrix)-1, len(matrix)*len(matrix[0])
	for {
		// 向左
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		// 向下
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		// 向右
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		// 向上
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++

		if len(ans) >= all {
			break
		}
	}
	return ans[:all]
}
