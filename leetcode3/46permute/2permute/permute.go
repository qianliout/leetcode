package main

func main() {

}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	dfs(nums, used, path, &ans)
	return ans
}

func dfs(nums []int, used []bool, path []int, ans *[][]int) {
	if len(path) == len(nums) {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	for i, ch := range nums {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, ch)
		dfs(nums, used, path, ans)
		used[i] = false
		path = path[:len(path)-1]
	}
}
