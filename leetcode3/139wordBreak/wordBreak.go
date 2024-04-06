package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

func wordBreak(s string, wordDict []string) bool {
	mem := make(map[string]bool)
	for i := range wordDict {
		mem[wordDict[i]] = true
	}
	dp := make([]bool, len(s)+1)
	ss := []byte(s)
	for i := 1; i <= len(s); i++ {
		if mem[string(ss[:i])] {
			dp[i] = true
		}
		for j := i - 1; j >= 0; j-- {
			if mem[string(ss[j:i])] {
				dp[i] = dp[i] || dp[j]
			}
		}
	}
	return dp[len(s)]
}

func wordBreak2(s string, wordDict []string) bool {
	mem := make(map[string]bool)
	for i := range wordDict {
		mem[wordDict[i]] = true
	}
	path := make([]string, 0)
	ans := false
	bfs([]byte(s), 0, mem, path, &ans)
	return ans
}

// timeout
func bfs(ss []byte, start int, wordDict map[string]bool, path []string, ans *bool) {
	if *ans {
		return
	}
	if start >= len(ss) && strings.Join(path, "") == string(ss) {
		*ans = true
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
