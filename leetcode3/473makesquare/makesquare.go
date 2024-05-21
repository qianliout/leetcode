package main

import (
	"fmt"
)

func main() {
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
	fmt.Println(makesquare([]int{3, 3, 3, 3, 4}))
	fmt.Println(makesquare([]int{5, 5, 5, 5, 16, 4, 4, 4, 4, 4, 3, 3, 3, 3, 4}))
}

func makesquare(matchsticks []int) bool {
	used := make([]bool, len(matchsticks))
	sm := 0
	for _, ch := range matchsticks {
		sm += ch
	}
	if sm%4 != 0 || len(matchsticks) == 0 {
		return false
	}
	find := dfs(matchsticks, used, 0, 0, 0, sm/4)
	return find
}

/*
k 表示已组装了多少个边了，cur表示当前组装的边的长，target，表示边长
*/
func dfs(matchs []int, used []bool, start, k, cur, target int) bool {
	if k == 4 {
		return true
	}

	if cur == target {
		return dfs(matchs, used, 0, k+1, 0, target)
	}

	for i := start; i < len(matchs); i++ {
		if used[i] {
			continue
		}
		if cur+matchs[i] > target {
			continue
		}
		used[i] = true
		b := dfs(matchs, used, i+1, k, cur+matchs[i], target)
		used[i] = false
		// 一个好的习惯是，在返回前把 user[i] = false
		if b {
			return true
		}
	}

	return false
}
