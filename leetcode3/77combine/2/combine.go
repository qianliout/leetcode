package main

func main() {

}

func combine(n int, k int) [][]int {
	path := make([]int, 0)
	ans := make([][]int, 0)
	dfs(n, k, 1, path, &ans)
	return ans

}

func dfs(n, k int, start int, path []int, ans *[][]int) {
	if len(path) == k {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	if len(path) > k {
		return
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		dfs(n, k, i+1, path, ans)
		path = path[:len(path)-1]
	}
}
