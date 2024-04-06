package main

import (
	"fmt"
)

func main() {
	fmt.Println(partition("aab"))
}

func partition(s string) [][]string {
	ss := []byte(s)
	path := make([]string, 0)
	mem := make(map[string]bool)
	ans := make([][]string, 0)
	dfs(ss, 0, path, &ans, mem)
	return ans
}

func dfs(ss []byte, start int, path []string, ans *[][]string, mem map[string]bool) {
	if start >= len(ss) {
		if len(path) > 0 {
			*ans = append(*ans, append([]string{}, path...))
			return
		}
	}
	for i := start + 1; i <= len(ss); i++ {
		pa := string(ss[start:i])
		if !isPalindrome(pa, mem) {
			continue
		}
		path = append(path, pa)
		dfs(ss, i, path, ans, mem)
		path = path[:len(path)-1]
	}
}

func isPalindrome(s string, mem map[string]bool) bool {
	if ss, ok := mem[s]; ok {
		return ss
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			mem[s] = false
			return false
		}
		left++
		right--
	}
	mem[s] = true
	return true
}
