package main

func main() {

}
func numIslands(grid [][]byte) int {
	ans := 0
	for i := range grid {
		for j, ch := range grid[i] {
			if ch == '1' {
				ans += 1
				full(grid, i, j)
			}
		}
	}
	return ans
}

func full(grid [][]byte, col, row int) {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return
	}
	if grid[col][row] == '0' {
		return
	}
	grid[col][row] = '0'
	full(grid, col+1, row)
	full(grid, col-1, row)
	full(grid, col, row+1)
	full(grid, col, row-1)
}
