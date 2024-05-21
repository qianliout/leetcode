package main

func main() {

}

func combinationSum3(k int, n int) [][]int {
	path := make([]int, 0)
	ans := make([][]int, 0)
	dfs(k, 9, 1, n, path, &ans)
	return ans
}

func dfs(k, n int, start, target int, path []int, ans *[][]int) {
	if target == 0 && len(path) == k {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	if target < 0 {
		return
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		dfs(k, n, i+1, target-i, path, ans)
		path = path[:len(path)-1]

	}
}
