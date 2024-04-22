package main

import (
	"fmt"
)

func main() {
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
}

func canCross(stones []int) bool {
	mem := make(map[int]int)
	for i, ch := range stones {
		mem[ch] = i
	}

	// 根据题意，第一步是固定经过步长 1 到达第一块石子（下标为 1）
	if _, ok := mem[1]; !ok {
		return false
	}
	cha := make(map[string]bool)
	return dfs(stones, mem, cha, 1, 1)
}

func dfs(stones []int, mem map[int]int, cha map[string]bool, start, k int) bool {
	if start >= len(stones)-1 {
		return true
	}
	key := fmt.Sprintf("%d_%d", start, k)
	if cur, ok := cha[key]; ok {
		return cur
	}

	for j := -1; j <= 1; j++ {
		if k+j == 0 {
			continue
		}
		next := stones[start] + k + j
		if idx, ok := mem[next]; ok {
			cur := dfs(stones, mem, cha, idx, k+j)
			key2 := fmt.Sprintf("%d_%d", idx, k+j)
			cha[key2] = cur
			if cur {
				return true
			}
		}
	}
	cha[key] = false

	return false
}
