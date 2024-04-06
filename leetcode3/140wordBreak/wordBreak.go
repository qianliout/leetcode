package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

func wordBreak(s string, wordDict []string) []string {
	mem := make(map[string]bool)
	for i := range wordDict {
		mem[wordDict[i]] = true
	}
	path := make([]string, 0)
	ans := make([]string, 0)
	bfs([]byte(s), 0, mem, path, &ans)
	return ans
}

// timeout
func bfs(ss []byte, start int, wordDict map[string]bool, path []string, ans *[]string) {
	if start >= len(ss) && strings.Join(path, "") == string(ss) {
		*ans = append(*ans, strings.Join(path, " "))
		return
	}

	for i := start + 1; i <= len(ss); i++ {
		ch := string(ss[start:i])
		if wordDict[ch] {
			path = append(path, ch)
			bfs(ss, i, wordDict, path, ans)
			path = path[:len(path)-1]
		}
	}
}
